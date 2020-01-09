// Package cli and it's sub packages implements the command line tool for Hasura
// GraphQL Engine. The CLI operates on a directory, denoted by
// "ExecutionDirectory" in the "ExecutionContext" struct.
//
// The ExecutionContext is passed to all the subcommands so that a singleton
// context is available for the execution. Logger and Spinner comes from the same
// context.
package cli

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hasura/graphql-engine/cli/metadata/actions"
	"github.com/hasura/graphql-engine/cli/plugins/paths"
	"github.com/hasura/graphql-engine/cli/telemetry"
	"github.com/hasura/graphql-engine/cli/util"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/briandowns/spinner"
	"github.com/gofrs/uuid"
	"github.com/hasura/graphql-engine/cli/version"
	"github.com/mattn/go-colorable"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Other constants used in the package
const (
	// Name of the global configuration directory
	GlobalConfigDirName = ".hasura"
	// Name of the global configuration file
	GlobalConfigFileName = "config.json"

	// Name of the file to store last update check time
	LastUpdateCheckFileName = "last_update_check_at"
)

// String constants
const (
	StrTelemetryNotice = `Help us improve Hasura! The cli collects anonymized usage stats which
allow us to keep improving Hasura at warp speed. To opt-out or read more,
visit https://docs.hasura.io/1.0/graphql/manual/guides/telemetry.html
`
)

// Config has the config values required to contact the server.
type Config struct {
	// Endpoint for the GraphQL Engine
	Endpoint string
	// AdminSecret (optional) required to query the endpoint
	AdminSecret string

	MetadataDirectory string

	Action actions.ActionExecutionConfig

	ParsedEndpoint *url.URL
}

type rawConfig struct {
	// Endpoint for the GraphQL Engine
	Endpoint string `json:"endpoint"`
	// AccessKey (deprecated) (optional) Admin secret key required to query the endpoint
	AccessKey string `json:"access_key,omitempty"`
	// AdminSecret (optional) Admin secret required to query the endpoint
	AdminSecret string `json:"admin_secret,omitempty"`

	MetadataDirectory string `json:"metadata_directory"`

	Action actions.ActionExecutionConfig `json:"action"`

	ParsedEndpoint *url.URL `json:"-"`
}

func (r rawConfig) toConfig() Config {
	s := r.AdminSecret
	if s == "" {
		s = r.AccessKey
	}
	return Config{
		Endpoint:          r.Endpoint,
		AdminSecret:       s,
		ParsedEndpoint:    r.ParsedEndpoint,
		MetadataDirectory: r.MetadataDirectory,
		Action:            r.Action,
	}
}

func (s Config) toRawConfig() rawConfig {
	return rawConfig{
		Endpoint:          s.Endpoint,
		AccessKey:         "",
		AdminSecret:       s.AdminSecret,
		ParsedEndpoint:    s.ParsedEndpoint,
		MetadataDirectory: s.MetadataDirectory,
		Action:            s.Action,
	}
}

// MarshalJSON converts s to JSON
func (s Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.toRawConfig())
}

// UnmarshalJSON converts b to struct s
func (s Config) UnmarshalJSON(b []byte) error {
	var r rawConfig
	err := json.Unmarshal(b, &r)
	if err != nil {
		return errors.Wrap(err, "unmarshal error")
	}
	sc := r.toConfig()
	s.Endpoint = sc.Endpoint
	s.AdminSecret = sc.AdminSecret
	s.ParsedEndpoint = sc.ParsedEndpoint
	s.MetadataDirectory = sc.MetadataDirectory
	s.Action = sc.Action
	return nil
}

// ParseEndpoint ensures the endpoint is valid.
func (s *Config) ParseEndpoint() error {
	nurl, err := url.Parse(s.Endpoint)
	if err != nil {
		return err
	}
	s.ParsedEndpoint = nurl
	return nil
}

// ExecutionContext contains various contextual information required by the cli
// at various points of it's execution. Values are filled in by the
// initializers and passed on to each command. Commands can also fill in values
// to be used further down the line.
type ExecutionContext struct {
	// CMDName is the name of CMD (os.Args[0]). To be filled in later to
	// correctly render example strings etc.
	CMDName string

	// ID is a unique ID for this Execution
	ID string

	// ServerUUID is the unique ID for the server this execution is contacting.
	ServerUUID string

	// Spinner is the global spinner object used to show progress across the cli.
	Spinner *spinner.Spinner
	// Logger is the global logger object to print logs.
	Logger *logrus.Logger

	// ExecutionDirectory is the directory in which command is being executed.
	ExecutionDirectory string
	// MigrationDir is the name of directory where migrations are stored.
	MigrationDir string
	// MetadataDir is the name of directory where metadata files are stored.
	MetadataDir string
	// ConfigFile is the file where endpoint etc. are stored.
	ConfigFile string

	// Config is the configuration object storing the endpoint and admin secret
	// information after reading from config file or env var.
	Config *Config

	// GlobalConfigDir is the ~/.hasura-graphql directory to store configuration
	// globally.
	GlobalConfigDir string
	// GlobalConfigFile is the file inside GlobalConfigDir where values are
	// stored.
	GlobalConfigFile string

	// GlobalConfig holds all the configuration options.
	GlobalConfig *GlobalConfig

	// IsStableRelease indicates if the CLI release is stable or not.
	IsStableRelease bool
	// Version indicates the version object
	Version *version.Version

	// Viper indicates the viper object for the execution
	Viper *viper.Viper

	// LogLevel indicates the logrus default logging level
	LogLevel string

	// NoColor indicates if the outputs shouldn't be colorized
	NoColor bool

	// Telemetry collects the telemetry data throughout the execution
	Telemetry *telemetry.Data

	// LastUpdateCheckFile is the file where the timestamp of last update check is stored
	LastUpdateCheckFile string

	// SkipUpdateCheck will skip the auto update check if set to true
	SkipUpdateCheck bool

	// PluginsPath is the path used by the plugins
	PluginsPath paths.Paths
}

// NewExecutionContext returns a new instance of execution context
func NewExecutionContext() *ExecutionContext {
	ec := &ExecutionContext{}
	ec.Telemetry = telemetry.BuildEvent()
	ec.Telemetry.Version = version.BuildVersion
	return ec
}

// Prepare as the name suggests, prepares the ExecutionContext ec by
// initializing most of the variables to sensible defaults, if it is not already
// set.
func (ec *ExecutionContext) Prepare() error {
	// set the command name
	cmdName := os.Args[0]
	if len(cmdName) == 0 {
		cmdName = "hasura"
	}
	ec.CMDName = cmdName

	// set spinner
	ec.setupSpinner()

	// set logger
	ec.setupLogger()

	// populate version
	ec.setVersion()

	// setup global config
	err := ec.setupGlobalConfig()
	if err != nil {
		return errors.Wrap(err, "setting up global config failed")
	}

	// setup plugins path
	err = ec.GetPluginsPath()
	if err != nil {
		return errors.Wrap(err, "setting up plugins path failed")
	}

	ec.LastUpdateCheckFile = filepath.Join(ec.GlobalConfigDir, LastUpdateCheckFileName)

	// initialize a blank server config
	if ec.Config == nil {
		ec.Config = &Config{}
	}

	// generate an execution id
	if ec.ID == "" {
		id := "00000000-0000-0000-0000-000000000000"
		u, err := uuid.NewV4()
		if err == nil {
			id = u.String()
		} else {
			ec.Logger.Debugf("generating uuid for execution ID failed, %v", err)
		}
		ec.ID = id
		ec.Logger.Debugf("execution id: %v", ec.ID)
	}
	ec.Telemetry.ExecutionID = ec.ID

	return nil
}

// GetPluginsPath returns the inferred paths for hasura. By default, it assumes
// $HOME/.hasura as the base path
func (ec *ExecutionContext) GetPluginsPath() error {
	home, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "cannot get home directory")
	}
	base := filepath.Join(home, GlobalConfigDirName, "plugins")
	base, err = filepath.Abs(base)
	if err != nil {
		return errors.Wrap(err, "cannot get absolute path")
	}
	ec.PluginsPath = paths.NewPaths(base)
	return nil
}

// Validate prepares the ExecutionContext ec and then validates the
// ExecutionDirectory to see if all the required files and directories are in
// place.
func (ec *ExecutionContext) Validate() error {

	// validate execution directory
	err := ec.validateDirectory()
	if err != nil {
		return errors.Wrap(err, "validating current directory failed")
	}

	// set names of files and directories
	ec.MigrationDir = filepath.Join(ec.ExecutionDirectory, "migrations")
	ec.ConfigFile = filepath.Join(ec.ExecutionDirectory, "config.yaml")

	// read config and parse the values into Config
	err = ec.readConfig()
	if err != nil {
		return errors.Wrap(err, "cannot read config")
	}

	if ec.Config.MetadataDirectory != "" {
		ec.MetadataDir = filepath.Join(ec.ExecutionDirectory, ec.Config.MetadataDirectory)
	}

	if _, err := os.Stat(ec.MetadataDir); os.IsNotExist(err) {
		err = os.MkdirAll(ec.MetadataDir, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "cannot write metadata directory")
		}
	}

	ec.Logger.Debug("graphql engine endpoint: ", ec.Config.Endpoint)
	ec.Logger.Debug("graphql engine admin_secret: ", ec.Config.AdminSecret)

	// get version from the server and match with the cli version
	err = ec.checkServerVersion()
	if err != nil {
		return errors.Wrap(err, "version check")
	}

	state := util.GetServerState(ec.Config.Endpoint, ec.Config.AdminSecret, ec.Version.ServerSemver, ec.Logger)
	ec.ServerUUID = state.UUID
	ec.Telemetry.ServerUUID = ec.ServerUUID
	ec.Logger.Debugf("server: uuid: %s", ec.ServerUUID)

	return nil
}

func (ec *ExecutionContext) checkServerVersion() error {
	v, err := version.FetchServerVersion(ec.Config.Endpoint)
	if err != nil {
		return errors.Wrap(err, "failed to get version from server")
	}
	ec.Version.SetServerVersion(v)
	ec.Telemetry.ServerVersion = ec.Version.GetServerVersion()
	isCompatible, reason := ec.Version.CheckCLIServerCompatibility()
	ec.Logger.Debugf("versions: cli: [%s] server: [%s]", ec.Version.GetCLIVersion(), ec.Version.GetServerVersion())
	ec.Logger.Debugf("compatibility check: [%v] %v", isCompatible, reason)
	if !isCompatible {
		return errors.Errorf("[cli: %s] [server: %s] versions incompatible: %s", ec.Version.GetCLIVersion(), ec.Version.GetServerVersion(), reason)
	}
	return nil
}

// readConfig reads the configuration from config file, flags and env vars,
// through viper.
func (ec *ExecutionContext) readConfig() error {
	// need to get existing viper because https://github.com/spf13/viper/issues/233
	v := ec.Viper
	v.SetEnvPrefix("HASURA_GRAPHQL")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigName("config")
	v.SetDefault("endpoint", "http://localhost:8080")
	v.SetDefault("admin_secret", "")
	v.SetDefault("access_key", "")
	v.SetDefault("metadata_directory", "")
	v.SetDefault("action.default_kind", "synchronous")
	v.SetDefault("action.default_handler", "{{DEFAULT_HANDLER}}")
	v.SetDefault("action.scaffold.output_dir", ec.ExecutionDirectory)
	v.AddConfigPath(ec.ExecutionDirectory)
	err := v.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "cannot read config from file/env")
	}
	adminSecret := v.GetString("admin_secret")
	if adminSecret == "" {
		adminSecret = v.GetString("access_key")
	}
	ec.Config = &Config{
		Endpoint:          v.GetString("endpoint"),
		AdminSecret:       adminSecret,
		MetadataDirectory: v.GetString("metadata_directory"),
		Action: actions.ActionExecutionConfig{
			Kind:    v.GetString("action.kind"),
			HandlerWebhookBaseURL: v.GetString("action.handler_webhook_baseurl"),
			Codegen: actions.CodegenExecutionConfig{
				Framework:         v.GetString("action.codegen.framework"),
				OutputDir:         v.GetString("action.codegen.output_dir"),
			},
		},
	}
	return ec.Config.ParseEndpoint()
}

// setupSpinner creates a default spinner if the context does not already have
// one.
func (ec *ExecutionContext) setupSpinner() {
	if ec.Spinner == nil {
		spnr := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
		spnr.Writer = os.Stderr
		ec.Spinner = spnr
	}
}

// Spin stops any existing spinner and starts a new one with the given message.
func (ec *ExecutionContext) Spin(message string) {
	ec.Spinner.Stop()
	ec.Spinner.Prefix = message
	ec.Spinner.Start()
}

// setupLogger creates a default logger if context does not have one set.
func (ec *ExecutionContext) setupLogger() {
	if ec.Logger == nil {
		logger := logrus.New()

		logger.Formatter = &logrus.TextFormatter{
			ForceColors:      true,
			DisableTimestamp: true,
		}
		logger.Out = colorable.NewColorableStdout()
		ec.Logger = logger
	}

	if ec.NoColor {
		ec.Logger.Formatter = &logrus.TextFormatter{
			DisableColors:    true,
			DisableTimestamp: true,
		}
	}

	if ec.LogLevel != "" {
		level, err := logrus.ParseLevel(ec.LogLevel)
		if err != nil {
			ec.Logger.WithError(err).Error("error parsing log-level flag")
			return
		}
		ec.Logger.SetLevel(level)
	}

	// set the logger for telemetry
	if ec.Telemetry.Logger == nil {
		ec.Telemetry.Logger = ec.Logger
	}
}

// SetVersion sets the version inside context, according to the variable
// 'version' set during build context.
func (ec *ExecutionContext) setVersion() {
	if ec.Version == nil {
		ec.Version = version.New()
	}
}

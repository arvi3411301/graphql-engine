package commands

import (
	"fmt"
	"strings"

	"github.com/hasura/graphql-engine/cli"
	"github.com/hasura/graphql-engine/cli/metadata/actions"
	"github.com/hasura/graphql-engine/cli/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newActionsCreateCmd(ec *cli.ExecutionContext) *cobra.Command {
	v := viper.New()
	opts := &actionsCreateOptions{
		EC: ec,
	}
	actionsCreateCmd := &cobra.Command{
		Use:          "create",
		Short:        "",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Viper = v
			err := ec.Prepare()
			if err != nil {
				return err
			}
			err = ec.Validate()
			if err != nil {
				return err
			}
			if ec.Config.Version != cli.V2 {
				return fmt.Errorf("actions commands can be executed only when config version is greater than 1")
			}
			if ec.MetadataDir == "" {
				return fmt.Errorf("actions commands can be executed only when metadata_dir is set in config")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.run()
		},
	}

	f := actionsCreateCmd.Flags()

	f.StringVar(&opts.deriveFrom, "derive-from", "", "derive action from a Hasura operation")
	f.BoolVar(&opts.withCodegen, "with-codegen", false, "create action along with codegen")

	f.String("endpoint", "", "http(s) endpoint for Hasura GraphQL Engine")
	f.String("admin-secret", "", "admin secret for Hasura GraphQL Engine")
	f.String("access-key", "", "access key for Hasura GraphQL Engine")
	f.MarkDeprecated("access-key", "use --admin-secret instead")
	f.String("kind", "", "kind to use in action")
	f.String("webhook", "", "webhook to use in action")

	// need to create a new viper because https://github.com/spf13/viper/issues/233
	v.BindPFlag("endpoint", f.Lookup("endpoint"))
	v.BindPFlag("admin_secret", f.Lookup("admin-secret"))
	v.BindPFlag("access_key", f.Lookup("access-key"))
	v.BindPFlag("actions.kind", f.Lookup("kind"))
	v.BindPFlag("actions.handler_webhook_baseurl", f.Lookup("webhook"))

	return actionsCreateCmd
}

type actionsCreateOptions struct {
	EC *cli.ExecutionContext

	name        string
	deriveFrom  string
	withCodegen bool
}

func (o *actionsCreateOptions) run() error {
	migrateDrv, err := newMigrate(o.EC, true)
	if err != nil {
		return err
	}

	// introspect Hasura schema if a mutation is being derived
	var introSchema interface{}
	if o.deriveFrom != "" {
		o.deriveFrom = strings.TrimSpace(o.deriveFrom)
		o.EC.Spin("Deriving a Hasura operation...")
		introSchema, err = migrateDrv.GetIntroSpectionSchema()
		if err != nil {
			return errors.Wrap(err, "error in fetching introspection schema")
		}
		o.EC.Spinner.Stop()
	}

	// create new action
	o.EC.Spin("Creating the action...")
	o.EC.Spinner.Stop()
	actionCfg := actions.New(o.EC, o.EC.MetadataDir)
	err = actionCfg.Create(o.name, introSchema, o.deriveFrom)
	if err != nil {
		return errors.Wrap(err, "error in creating action")
	}
	err = migrateDrv.ApplyMetadata()
	if err != nil {
		return errors.Wrap(err, "error in applying metadata")
	}

	o.EC.Spinner.Stop()
	o.EC.Logger.WithField("name", o.name).Infoln("action created")

	// if codegen config not present, skip codegen
	if o.EC.Config.ActionConfig.Codegen.Framework == "" {
		if o.withCodegen {
			return fmt.Errorf(`Could not find codegen config. For adding codegen config, run:

  hasura actions use-codegen`)
		}
		return nil
	}

	// if with-codegen flag not present, ask them if they want to codegen
	var confirmation string
	if !o.withCodegen {
		confirmation, err = util.GetYesNoPrompt("Do you want to generate " + o.EC.Config.ActionConfig.Codegen.Framework + " code for this action and the custom types?")
		if err != nil {
			return errors.Wrap(err, "error in getting user input")
		}
	}

	if confirmation == "n" {
		infoMsg := fmt.Sprintf(`You skipped codegen. For getting codegen for this action, run:

  hasura actions codegen %s
`, o.name)
		o.EC.Logger.Info(infoMsg)
		return nil
	}

	// construct derive payload to send to codegenerator
	derivePayload := actions.DerivePayload{
		IntrospectionSchema: introSchema,
		Operation:           o.deriveFrom,
		ActionName:          o.name,
	}

	// Run codegen
	o.EC.Spin(fmt.Sprintf(`Running "hasura actions codegen %s"...`, o.name))
	err = actionCfg.Codegen(o.name, derivePayload)
	if err != nil {
		return errors.Wrap(err, "error in generating codegen")
	}
	o.EC.Spinner.Stop()
	o.EC.Logger.Info("Codegen files generated at " + o.EC.Config.ActionConfig.Codegen.OutputDir)
	return nil

}

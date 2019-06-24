package commands

import (
	"github.com/hasura/graphql-engine/cli"
	"github.com/hasura/graphql-engine/cli/metadata"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const longHelpMetadataExportCmd = `Export Hasura metadata and save it in migrations/metadata.yaml file.
The output is a yaml file which captures all the metadata required
by GraphQL Engine. This includes info about tables that are tracked,
permission rules, relationships and event triggers that are defined
on those tables.`

func newMetadataExportCmd(ec *cli.ExecutionContext) *cobra.Command {
	v := viper.New()
	opts := &metadataExportOptions{
		EC: ec,
	}

	metadataExportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export Hasura GraphQL Engine metadata from the database",
		Example: `  # Export metadata and save it in migrations/metadata.yaml file:
  hasura metadata export`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Viper = v
			return ec.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.EC.Spin("Exporting metadata...")
			err := opts.run()
			opts.EC.Spinner.Stop()
			if err != nil {
				return errors.Wrap(err, "failed to export metadata")
			}
			opts.EC.Logger.Info("Metadata exported")
			return nil
		},
		Long: longHelpMetadataExportCmd,
	}

	f := metadataExportCmd.Flags()
	f.String("endpoint", "", "http(s) endpoint for Hasura GraphQL Engine")
	f.String("admin-secret", "", "admin secret for Hasura GraphQL Engine")
	f.String("access-key", "", "access key for Hasura GraphQL Engine")
	f.MarkDeprecated("access-key", "use --admin-secret instead")

	// need to create a new viper because https://github.com/spf13/viper/issues/233
	v.BindPFlag("endpoint", f.Lookup("endpoint"))
	v.BindPFlag("admin_secret", f.Lookup("admin-secret"))
	v.BindPFlag("access_key", f.Lookup("access-key"))

	return metadataExportCmd
}

type metadataExportOptions struct {
	EC *cli.ExecutionContext
}

func (o *metadataExportOptions) run() error {
	config, err := metadata.New(o.EC.MigrationDir, o.EC.ServerConfig.ParsedEndpoint, o.EC.ServerConfig.AdminSecret, o.EC.Version)
	if err != nil {
		return err
	}

	metadataPath, err := o.EC.GetMetadataFilePath("yaml")
	if err != nil {
		return errors.Wrap(err, "cannot save metadata")
	}

	err = config.SetMetadataPath(metadataPath, false)
	if err != nil {
		return errors.Wrap(err, "cannot save metadata")
	}
	return config.Export()
}

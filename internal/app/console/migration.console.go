package console

import (
	"github.com/bagastri07/gin-boilerplate-service/internal/app/bootstrap"
	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command.
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Performs database migration operations",
	Long:  `The migration command is used to perform database migration operations, such as applying or rolling back migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		action, _ := cmd.Flags().GetString("action")
		migrationName, _ := cmd.Flags().GetString("name")
		version, _ := cmd.Flags().GetInt64("version")

		bootstrap.StartMigration(action, migrationName, &version)
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
	migrationCmd.PersistentFlags().String("action", "up", "action create|up|up-by-one|up-to|down|down-to|reset|status")
	migrationCmd.PersistentFlags().Int64("version", 1, "version")
	migrationCmd.PersistentFlags().String("name", "", "migration name")
}

package cmd

import (
	"github.com/spf13/cobra"
)

var migrationCMD = &cobra.Command{
	Use:   "migration",
	Short: "Run migration",
	Run:   runMigration,
}

func init() {
	rootCMD.AddCommand(migrationCMD)
}

func runMigration(c *cobra.Command, args []string) {
}

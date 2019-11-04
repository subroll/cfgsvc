package cmd

import (
	"github.com/spf13/cobra"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Run migration",
	Run:   runMigration,
}

func init() {
	rootCMD.AddCommand(migrateCMD)
}

func runMigration(c *cobra.Command, args []string) {
}

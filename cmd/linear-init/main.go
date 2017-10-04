package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ereOn/linear/pkg/command"
	"github.com/ereOn/linear/pkg/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear init",
	Short: "Start a new project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		path, err := filepath.Abs(database.DefaultDatabaseFileName)

		if err != nil {
			return err
		}

		_, err = database.FromFile(path)

		if err == nil {
			return fmt.Errorf("a database already exists at `%s`", path)
		}

		db := database.Database{}

		return db.ToFile(path)
	},
}

func main() {
	command.ImplementDescribe(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

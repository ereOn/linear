package main

import (
	"os"

	"github.com/ereOn/linear/pkg/command"
	"github.com/ereOn/linear/pkg/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear add api <name>",
	Short: "Add a new API to the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cobra.NoArgs(cmd, args)
		}

		name := args[0]

		cmd.SilenceUsage = true

		scanner, err := database.NewDefaultScanner()

		if err != nil {
			return err
		}

		path, db, err := scanner.Scan()

		if err != nil {
			return err
		}

		component := database.Component{
			ID: database.ComponentID{
				Type: "api",
				Name: name,
			},
		}

		if err := db.Add(component); err != nil {
			return err
		}

		return db.ToFile(path)
	},
}

func main() {
	command.ImplementDescribe(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

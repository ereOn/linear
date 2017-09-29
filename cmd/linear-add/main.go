package main

import (
	"fmt"
	"os"

	"github.com/ereOn/linear/pkg/command"
	"github.com/ereOn/linear/pkg/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear add",
	Short: "Add a new component to an existing project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		scanner, err := database.NewDefaultScanner()

		if err != nil {
			return err
		}

		path, db, err := scanner.Scan()

		if err != nil {
			return err
		}

		fmt.Println(path, db)

		return nil
	},
}

func main() {
	command.ImplementDescribe(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/ereOn/linear/pkg/command"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear project",
	Short: "Start a new project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	command.ImplementDescribe(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/ereOn/linear/pkg/command"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear",
	Short: "A project bootstraper with style",
}

func main() {
	scanner := command.Scanner{
		CommandPrefix: "linear",
	}
	commands, errors := scanner.Scan()

	for _, command := range commands {
		rootCmd.AddCommand(command.AsCobraCommand())
	}

	if len(errors) > 0 {
		fmt.Fprintf(os.Stderr, "Errors while scanning for commands:\n")

		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "- %s\n", err)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linear",
	Short: "A project bootstraper with style",
}

var (
	subcommandsBinaries = getSubcommandBinaries(os.Getenv("PATH"))
	subcommands         = getSubcommands(subcommandsBinaries)
)

func main() {
	rootCmd.AddCommand(subcommands...)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

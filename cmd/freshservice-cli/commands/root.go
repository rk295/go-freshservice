package commands

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "freshservice-cli",
		Short: "A simple Freshservice command line interface",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

package cmd

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "base",
	Short: "Example base command",
}

func Execute() {
	cobra.CheckErr(RootCommand.Execte())
}

func init() {
	// nothing to do
}

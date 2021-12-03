package cmd

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "extension",
	Short: "Example extended command",
}

func Execute() {
	cobra.CheckErr(RootCommand.Execute())
}

func init() {
	// nothing to do
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Command1 = &cobra.Command {
	Use: "command-1",
	Short: "Example command: command-1",
	RunE: func(cmd *cobra.Command) error {
		fmt.Println ("base command-1")
		return nil
	}
}

func init() {
	RootCommand.AddCommand(Command1)
}
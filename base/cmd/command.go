package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "command",
	Short: "Example command: command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("base command")
		return nil
	},
}

func init() {
	RootCommand.AddCommand(Command)
}

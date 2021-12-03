package cmd

import (
	"fmt"

	base "github.com/kalantar/testcli/base/cmd"
	"github.com/spf13/cobra"
)

var Command = base.Command

func init() {
	Command.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("extension command")
		return nil
	}
	RootCommand.AddCommand(Command)
}

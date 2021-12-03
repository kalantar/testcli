package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NewCommand = &cobra.Command{
	Use:   "new",
	Short: "Example command: new",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("extension new")
		return nil
	},
}

func init() {
	RootCommand.AddCommand(NewCommand)
}

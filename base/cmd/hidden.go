package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Hidden = &cobra.Command{
	Use:   "hidden",
	Short: "Example command: hidden",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("base hidden")
		return nil
	},
}

func init() {
	RootCommand.AddCommand(Hidden)
}

package cmd

import (
)

func NewCmdIter8Remote(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewCmdIter8RemoteOptions(streams)

	cmd := &cobra.Command{
		Use: "iter8r runr [flags]",
		Short: "Run an experiment in the cluster",
		Example: "Run an experiment in the cluster",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

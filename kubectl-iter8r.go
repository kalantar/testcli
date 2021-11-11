/*
*/

package main

import (
	"os"

	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"github.com/kalantar/iter8r/pkg/cmd"
)

func main() {
//	flags := pflag.NewFlagSet("kubectl-iter8r", pflag.ExitOnError)
//	pflag.CommandLine = flags

	root := cmd.NewCmdIter8Remote(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

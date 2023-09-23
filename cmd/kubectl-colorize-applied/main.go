package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cappyzawa/kubectl-colorize-applied/internal/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type cli struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

func (c *cli) run(ctx context.Context) int {
	flags := pflag.NewFlagSet("kubectl-ns", pflag.ExitOnError)
	pflag.CommandLine = flags

	streams := genericclioptions.IOStreams{In: c.in, Out: c.out, ErrOut: c.err}
	root := cmd.NewCmdColorizeApplied(ctx, streams)
	if err := root.Execute(); err != nil {
		fmt.Fprintf(c.err, "error: %v\n", err)
		return 1
	}
	return 0
}

func main() {
	c := &cli{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
	ctx := context.Background()
	os.Exit(c.run(ctx))
}

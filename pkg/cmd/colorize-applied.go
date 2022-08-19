package cmd

import (
	"bufio"
	"context"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type ColorizeAppliedOptions struct {
	configFlags *genericclioptions.ConfigFlags
	args        []string

	genericclioptions.IOStreams
}

func NewCmdColorizeAppliedOptions(streams genericclioptions.IOStreams) *ColorizeAppliedOptions {
	return &ColorizeAppliedOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

func NewCmdColorizeApplied(ctx context.Context, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewCmdColorizeAppliedOptions(streams)

	cmd := &cobra.Command{
		Use:          "colorize-applied [flags]",
		Short:        "colorize applied result",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run(ctx)
		},
	}
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for processing the command.
func (o *ColorizeAppliedOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args
	return nil
}

// Validate ensures that all required arguments and flag values are provided.
func (o *ColorizeAppliedOptions) Validate() error {
	return nil
}

// Run runs the command.
func (o *ColorizeAppliedOptions) Run(ctx context.Context) error {
	scanner := bufio.NewScanner(o.In)
	for scanner.Scan() {
		o.ColorPrint(scanner.Text())
	}
	return nil
}

func (o *ColorizeAppliedOptions) ColorPrint(line string) (int, error) {
	// key is keyword
	// value is color attributes
	m := map[string][]color.Attribute{
		"pruned":     {color.FgRed},
		"configured": {color.FgYellow},
		"created":    {color.FgGreen},
		"unchanged":  {color.FgHiBlack},
	}
	for k, v := range m {
		if strings.Contains(line, k) {
			return color.New(v...).Fprintf(o.Out, "%s\n", line)
		}
	}
	return o.Out.Write([]byte(line + "\n"))
}

package cmd_test

import (
	"bytes"
	"testing"

	"github.com/cappyzawa/kubectl-colorize-applied/internal/cmd"
	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestColorizeAppliedOptionsColorPrint(t *testing.T) {
	color.NoColor = false
	t.Parallel()

	cases := map[string]struct {
		input  string
		output string
	}{
		"pruned": {
			input:  "resource pruned",
			output: color.RedString("resource pruned\n"),
		},
		"configured": {
			input:  "resource configured",
			output: color.YellowString("resource configured\n"),
		},
		"created": {
			input:  "resource created",
			output: color.GreenString("resource created\n"),
		},
		"unchanged": {
			input:  "resource unchanged",
			output: color.HiBlackString("resource unchanged\n"),
		},
		"resource name contains unchanged": {
			input:  "unchanged-resource pruned",
			output: color.RedString("unchanged-resource pruned\n"),
		},
	}
	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := bytes.NewBuffer([]byte{})
			o := cmd.NewCmdColorizeAppliedOptions(genericclioptions.IOStreams{Out: actual})
			o.ColorPrint(tt.input)
			if diff := cmp.Diff(tt.output, actual.String()); diff != "" {
				t.Errorf("unexpected output (-want, +got):\n%s", diff)
			}
		})
	}
}

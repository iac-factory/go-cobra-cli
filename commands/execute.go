package commands

import (
	"github.com/iac-factory/go-cobra-cli/commands/example"
	"github.com/iac-factory/go-cobra-cli/internal/log/color"
	"github.com/spf13/cobra"
)

// Execute runs the root command and handles any CLI execution exception. Additionally,
// all child command(s) are added to the root command.
func Execute(root *cobra.Command) {
	root.AddCommand(example.Command)

	if e := root.Execute(); e != nil {
		color.New().Bold(
			color.New().Red("error").String(),
		).Default("-").Italic(
			color.New().White(e.Error()).String(),
		).Write()
	}
}

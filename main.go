package main

import (
	_ "embed"
	"fmt"
	"github.com/iac-factory/go-cobra-cli/internal/commands"
	"github.com/iac-factory/go-cobra-cli/internal/constants"
	"github.com/iac-factory/go-cobra-cli/internal/log"
	"github.com/iac-factory/go-cobra-cli/internal/types/level"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strings"
)

var version string = "0.0.0" // see go linking for compile-time variable overwrites

// logging is a variable that represents the current log level configuration.
var (
	logging level.Type = level.Error
)

func main() {
	var root = &cobra.Command{
		Use:                    constants.Name(),
		Short:                  "A cli tool [...]",
		Long:                   "A cli tool [...]",
		Example:                "",
		ValidArgs:              nil,
		ValidArgsFunction:      nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Annotations:            nil,
		Version:                version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			var l level.Type

			if e := l.Set(logging.String()); e != nil {
				return e
			} else {
				log.Default(l.String())
			}

			slog.Log(ctx, log.Trace, "Root", slog.Group("command",
				slog.String("name", cmd.Name()),
				slog.String("version", version),
				slog.Group("flags",
					slog.String("log-level", logging.String()),
				),
				slog.Group("environment",
					slog.String("LOG_LEVEL", os.Getenv("LOG_LEVEL")),
				),
			))

			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// @todo Logic to check if a newer version is available
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				if e := cmd.Help(); e != nil {
					panic(e)
				}
			}
		},
		PostRun:           nil,
		CompletionOptions: cobra.CompletionOptions{},
		TraverseChildren:  true,
		Hidden:            false,
		SilenceErrors:     false,
		SilenceUsage:      false,
	}

	root.PersistentFlags().VarP(&logging, "verbosity", "v", "sets and configures logging verbosity")

	commands.Execute(root)
}

func init() {
	version = strings.TrimSpace(version)
	if e := os.Setenv("VERSION", version); e != nil {
		exception := fmt.Errorf("unable to set VERSION: %w", e)

		panic(exception)
	}
}

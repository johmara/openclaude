package cmd

import (
	"fmt"
	"os"

	"github.com/johmara/openclaude/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openclaude",
	Short: "A full-featured TUI for Claude Code",
	Long:  "OpenClaude wraps Claude Code CLI in a rich terminal interface with streaming markdown, tool call visualization, and modal dialogs.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

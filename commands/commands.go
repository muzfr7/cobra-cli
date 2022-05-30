package commands

import (
	"github.com/spf13/cobra"
)

// init initialises all commands, add your commands here for initialization
// example: RootCmd.AddCommand(Cmd1, Cmd2, Cmd3).
func init() {
	RootCmd.AddCommand(Cmd1)
}

// RootCmd root command.
var RootCmd = &cobra.Command{
	Use:           "cobra-cli",
	Short:         "cobra-cli – demonstration command-line tool",
	SilenceErrors: true,
	SilenceUsage:  true,
}

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// init initialises all commands, add your commands here for initialization
// example: RootCmd.AddCommand(cmd1, cmd2, cmd3).
func init() {
	RootCmd.AddCommand(cmd1)
}

// RootCmd root command.
var RootCmd = &cobra.Command{
	Use:           "cli",
	Short:         "cli – demonstration command-line tool",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// cmd1 sample command.
var cmd1 = &cobra.Command{
	Use:   "cmd1",
	Short: "cmd1 does something",
	Run: func(ccmd *cobra.Command, args []string) {
		fmt.Println("Hello from cmd one")
	},
}

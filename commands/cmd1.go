package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd1 sample command.
var Cmd1 = &cobra.Command{
	Use:   "cmd1",
	Short: "cmd1 does something",
	Run:   cmd1,
}

func cmd1(ccmd *cobra.Command, args []string) {
	fmt.Println("Hello from cmd one")
}

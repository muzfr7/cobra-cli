package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmd1 sample command.
var cmd1 = &cobra.Command{
	Use:   "cmd1",
	Short: "cmd1 does something",
	Run:   cmdOne,
}

func cmdOne(ccmd *cobra.Command, args []string) {
	fmt.Println("Hello from cmd one")
}

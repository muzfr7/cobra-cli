package main

import (
	"log"

	"github.com/muzfr7/cobra-cli/internal/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

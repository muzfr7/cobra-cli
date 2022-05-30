package main

import (
	"log"

	"github.com/muzfr7/cobra-cli/internal/command"
)

func main() {
	if err := command.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

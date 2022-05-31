package main

import (
	"log"

	"github.com/muzfr7/cobra-cli/config"
	"github.com/muzfr7/cobra-cli/internal/commands"
)

var (
	// Env will be transformed to store environment variables.
	Env config.EnvConfig
)

func init() {
	// load environment variables from .env file into Env.
	if err := Env.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

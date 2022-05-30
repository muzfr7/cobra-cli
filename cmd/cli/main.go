package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/muzfr7/cobra-cli/config"
	"github.com/muzfr7/cobra-cli/internal/commands"
)

var (
	// Env will be transformed to store environment variables.
	Env config.EnvConfig
)

func init() {
	// export environment variables from .env file.
	if _, err := os.Stat("./.env"); err == nil {
		if err = config.ExportEnvVars("./.env"); err != nil {
			log.Fatal(err)
		}
	}

	// load exported environment variables into Env struct.
	if err := envconfig.Process("", &Env); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

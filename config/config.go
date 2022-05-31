package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// EnvConfig holds all environment variables.
type EnvConfig struct {
	AppName     string `envconfig:"APP_NAME" required:"true"`
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
	LogLevel    string `envconfig:"LOG_LEVEL" required:"true"`
}

// Load loads environment variables into EnvConfig.
func (e *EnvConfig) Load() error {
	// check whether .env file is present.
	if _, err := os.Stat("./.env"); err != nil {
		return err
	}

	// export environment variables from .env file.
	if err := exportEnvVars("./.env"); err != nil {
		return err
	}

	// load exported environment variables into Env struct.
	if err := envconfig.Process("", e); err != nil {
		return err
	}

	return nil
}

// exportEnvVars exports environment variables from given .env file.
func exportEnvVars(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "export ") {
			line = line[7:]
		}

		token := strings.SplitN(line, "=", 2)

		// Remove spaces, ' and "
		key := strings.Trim(strings.Trim(strings.TrimSpace(token[0]), "'"), "\"")
		value := strings.Trim(strings.Trim(strings.TrimSpace(token[1]), "'"), "\"")

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"

	configs "github.com/dshmmd/quera-snapp-bootcamp-2024/config"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/internal/server"
)

func main() {
	// Define the config flag using pflag
	configPath := pflag.StringP("config", "c", "", "Path to the configuration YAML file")

	// Parse the command-line flags
	pflag.Parse()

	// Check if the config path is provided
	if *configPath == "" {
		log.Fatalln("configuration file path is required. Use -config or -c to specify it.")
	}

	// Load the configuration from the YAML file
	config, err := configs.NewConfigFromYaml(*configPath)
	if err != nil {
		log.Fatalln(fmt.Errorf("can not read config file: %w", err))
	}

	// Create the server instance using the loaded config
	srv, err := server.NewServer(config)
	if err != nil {
		log.Fatalln(fmt.Errorf("can not create server: %w", err))
	}

	// Start the server
	err = srv.Serve()
	if err != nil {
		log.Fatalln(fmt.Errorf("error while running server: %w", err))
	}
}

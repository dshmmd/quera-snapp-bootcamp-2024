package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// NewConfigFromYaml reads a YAML file and unmarshals it into a Config struct.
func NewConfigFromYaml(fileAddress string) (*Config, error) {
	// Open the YAML files
	file, err := os.Open(fileAddress)
	if err != nil {
		return nil, fmt.Errorf("error opening YAML file: %v\n", err)
	}
	defer file.Close()

	// Create an instance of Config to hold the unmarshalled data
	config := &Config{}

	// Decode the file contents into the config struct
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("error decoding YAML file: %v\n", err)
	}

	return config, nil
}

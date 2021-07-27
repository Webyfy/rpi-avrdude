package main

import (
	"encoding/json"
	"os"
)

// Config represents application configuration
type Config struct {
	DefaultResetPin int    `json:"defaultResetPin"`
	Ports           []Port `json:"ports"`
}

// Name defines a port and its reset Pin
type Port struct {
	Name     string `json:"name"`
	ResetPin int    `json:"resetPin"`
}

// loadConfiguration loads configuration file and converts it into Config type
func loadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config represents application configuration
type Config struct {
	DefaultPin int    `json:"defaultPin"`
	Ports      []Port `json:"ports"`
}

// Name defines a port and its reset Pin
type Port struct {
	Name     string `json:"name"`
	ResetPin int    `json:"resetPin"`
}

// loadConfiguration loads configuration file and converts it into Config type
func loadConfiguration(configFile string) (Config, error) {
	var config Config

	if filepath.IsAbs(configFile) {
		viper.SetConfigFile(configFile)
	} else {
		ex, err := os.Executable()
		if err != nil {
			return config, err
		}
		exPath := filepath.Dir(ex)
		viper.SetConfigName(strings.TrimSuffix(filepath.Base(configFile), filepath.Ext(configFile)))
		viper.AddConfigPath(filepath.Dir(configFile))
		viper.AddConfigPath(exPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

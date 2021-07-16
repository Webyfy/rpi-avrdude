package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config represents application configuration
type Config struct {
	DefaultPin int          `json:"defaultPin"`
	PortPinMap []PortPinMap `json:"portPinMap"`
}

// PortPinMap represents mapping from a serial port
// to its corresponding reset pin
type PortPinMap struct {
	Port string `json:"port"`
	Pin  int    `json:"pin"`
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

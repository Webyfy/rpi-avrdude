package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kataras/golog"
)

const (
	OriginalFile = "avrdude-original"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "Path to configuration file")
	flag.Parse()

	var err error
	config, err := loadConfiguration(configFile)
	if err != nil {
		golog.Fatal("Failed to load configuration file: %w", err)
	}
	golog.Info("Config:", fmt.Sprintf("%+v", config))

	ownDir := getOwnDir()
	originalExec := filepath.Join(ownDir, OriginalFile)
	golog.Info(originalExec)

}

func getOwnDir() string {
	exec, err := os.Executable()
	if err != nil {
		golog.Fatal(err)
	}
	return filepath.Dir(exec)
}

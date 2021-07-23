package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"git.reach-iot.com/iot-master/rpi-avrdude/gpio"
	"github.com/kataras/golog"
)

const (
	OriginalFile = "avrdude-original"
	resetPin     = 23
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

	// Blink
	pin := gpio.NewDigitalPin(resetPin)
	if err := pin.Export(); err != nil {
		log.Fatal(err)
	}
	defer pin.Unexport()
	pin.Direction(gpio.OUT)
	for i := 0; i < 30; i++ {
		pin.Write(gpio.HIGH)
		time.Sleep(time.Second)
		pin.Write(gpio.LOW)
		time.Sleep(time.Second)
	}

}

func getOwnDir() string {
	exec, err := os.Executable()
	if err != nil {
		golog.Fatal(err)
	}
	return filepath.Dir(exec)
}

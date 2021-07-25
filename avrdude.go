package main

import (
	"os"
	"os/exec"
	"time"

	"git.reach-iot.com/iot-master/rpi-avrdude/gpio"
)

func normalRun(exe string, args []string) {
	cmd := exec.Command(exe, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
}

func straceRun(exe string, args []string) {

}

func reset(pin int) error {
	digitalPin := gpio.NewDigitalPin(pin)
	err := digitalPin.Export()
	if err != nil {
		return err
	}
	defer digitalPin.Unexport()

	if err := digitalPin.Direction(gpio.OUT); err != nil {
		return err
	}
	if err := digitalPin.Write(gpio.LOW); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 250)
	if err := digitalPin.Write(gpio.HIGH); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 50)

	return nil

}

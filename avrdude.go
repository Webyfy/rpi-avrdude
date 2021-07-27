package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"git.reach-iot.com/iot-master/rpi-avrdude/gpio"
)

const (
	dtrRequestPattern = `.+TIOCM_DTR.+`
	exitStatusPattern = `.+exited with.+`
)

// avrdudeProxy runs original avrdude in both normal mode
// as well as GPIO reset mode
type avrdudeProxy struct {
	orignalExec string
	args        []string
	resetPin    int
}

// normalRun runs original avrdude executable without any modification
func (a avrdudeProxy) normalRun() {
	cmd := exec.Command(a.orignalExec, a.args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// gpioResetRun runs original avrdude executable with strace
func (a avrdudeProxy) gpioResetRun() {
	cmdArgs := []string{"-eioctl", a.orignalExec}
	cmdArgs = append(cmdArgs, a.args...)
	cmd := exec.Command("strace", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	cmdReader, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer cmdReader.Close()

	cmd.Start()
	err = a.watchOutput(cmdReader)
	if err != nil {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		log.Fatal(err)
	}
	cmd.Wait()
}

// watchOutput reads strace output and resets arduino
// via GPIO pin on DTR trigger system call
func (a avrdudeProxy) watchOutput(cmdReader io.Reader) error {
	// regex
	dtrRegex, err := regexp.Compile(dtrRequestPattern)
	if err != nil {
		return err
	}
	exitRegex, err := regexp.Compile(exitStatusPattern)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	done := false
	for scanner.Scan() {
		line := scanner.Bytes()
		if dtrRegex.Match(line) {
			if !done {
				log.Println("Using autoreset DTR on GPIO Pin ", a.resetPin)
				a.reset()
				done = true
			}
		} else if exitRegex.Match(line) {
			log.Println(string(line))
		}
	}

	return nil
}

// reset resets microcontroller via GPIO pin
// timings were copied from https://github.com/andygock/avrdude-arduino/blob/37e1f67e6622a676a3866f8acd6a9618551941ca/stk500v2.c#L1319
func (a avrdudeProxy) reset() error {
	pin := gpio.NewDigitalPin(a.resetPin)
	if err := pin.Export(); err != nil {
		return err
	}
	defer pin.Unexport()

	if err := pin.Direction(gpio.OUT); err != nil {
		return err
	}
	if err := pin.Write(gpio.LOW); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 250)
	if err := pin.Write(gpio.HIGH); err != nil {
		time.Sleep(time.Millisecond * 50)
	}

	return nil
}

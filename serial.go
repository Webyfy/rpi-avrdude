package main

import (
	"errors"

	"go.bug.st/serial/enumerator"
)

func isUSBConverter(destinationPort string) (bool, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return false, err
	}
	if len(ports) == 0 {
		return false, errors.New("no serial port found")
	}
	for _, port := range ports {
		if port.Name == destinationPort {
			return port.IsUSB, nil
		}
	}

	return false, errors.New("port not found")
}

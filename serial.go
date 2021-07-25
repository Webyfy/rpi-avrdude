package main

import (
	"errors"

	"go.bug.st/serial/enumerator"
)

var errPortNotFound = errors.New("port not found")

func isGpioUart(destinationPort string) (bool, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return false, err
	}
	for _, port := range ports {
		if port.Name == destinationPort {
			return !port.IsUSB, nil
		}
	}

	return false, errPortNotFound
}

build:
	go build -o ~/.arduino15/packages/arduino/tools/avrdude/6.3.0-arduino17/bin/avrdude
	cp -f config.json ~/.arduino15/packages/arduino/tools/avrdude/6.3.0-arduino17/bin/

.PHONY: build
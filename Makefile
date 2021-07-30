# GO=go
RM=rm -rf
# MD=mkdir -p
CP=cp -f

OS=linux
ARCH=arm
ARM_VERSIONS=5 6

BUILD=build
BIN_NAME=avrdude
RELEASE=$$(git tag -l)
CONFIG_NAME=config.json
ARCHIVE_PREFIX=rpi-avrdude
ARCHIVE_SUFFIX=.tar.gz

define build_pack

env GOOS=linux GOARCH=arm GOARM=$(1) go build -o $(BUILD)/$(1)/$(BIN_NAME)
$(CP) $(CONFIG_NAME) $(BUILD)/$(1)/
cd $(BUILD)/$(1); tar -czf ../$(ARCHIVE_PREFIX)_$(RELEASE)_$(OS)_$(ARCH)_$(1).tar.gz *

endef

default: clean all

all:
	$(foreach VERSION,$(ARM_VERSIONS),$(call build_pack,$(VERSION)))

clean:
	$(RM) $(BUILD)

.PHONY: default all clean

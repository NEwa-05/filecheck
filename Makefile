BIN_NAME := filecheck
MAIN_DIRECTORY := ./

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION ?= $(if $(TAG_NAME),$(TAG_NAME),v0.0.0-$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

DOCKER_BUILD_PLATFORMS ?= linux/amd64, windows/amd64

## Build targets:
./bin/%/$(BIN_NAME): binpath $(BUILD_DEPS)
	@echo SHA: $(SHA) $(BUILD_DATE)
	$(eval GOOS := $(call word-slash, $*, 1))
	$(eval GOARCH := $(call word-slash, $*, 2))
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_FLAGS) -o $@ ${MAIN_DIRECTORY}

binpath:
	mkdir -p bin

# Target for building development image.
.PHONY: build
build: $(MAKE) -C $(CURDIR) build

.PHONY: build-linux-amd64
build-linux-amd64: build ./dist/linux/amd64/$(BIN_NAME)

.PHONY: build-windows-amd64
build-windows-amd64: build ./dist/windows/amd64/$(BIN_NAME)

# Reset implicit suffixes and force go suffix only.
.SUFFIXES:
.SUFFIXES: .go

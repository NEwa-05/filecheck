BIN_NAME := filecheck

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION ?= $(if $(TAG_NAME),$(TAG_NAME),v0.0.0-$(SHA))
VERSION_GIT := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GOGC ?=

.PHONY: default
#? default: Run `make binary`
default: binary

binpath:
	mkdir -p bin

binary: binpath
	CGO_ENABLED=1 GOGC=${GOGC} GOOS=${GOOS} GOARCH=${GOARCH} go build \
	-ldflags="-extldflags=-Wl,-no_warn_duplicate_libraries" \
	-installsuffix nocgo -o "./bin/${GOOS}/${GOARCH}/$(BIN_NAME)"

binary-linux-amd64: export GOOS := linux
binary-linux-amd64: export GOARCH := amd64
binary-linux-amd64:
	@$(MAKE) binary

binary-windows-amd64: export GOOS := windows
binary-windows-amd64: export GOARCH := amd64
binary-windows-amd64: export BIN_NAME := filecheck.exe
binary-windows-amd64:
	@$(MAKE) binary

binary-darwin-arm64: export GOOS := darwin
binary-darwin-arm64: export GOARCH := arm64
binary-darwin-arm64:
	@$(MAKE) binary

SHELL := /bin/bash

BINARY_NAME := myapp
DIST_DIR := dist

OSES := linux windows darwin
ARCHES := amd64 arm64

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: all build clean

all: build

build:
	@mkdir -p $(DIST_DIR)
	@set -e; \
	for os in $(OSES); do \
		for arch in $(ARCHES); do \
			echo "▶ Building for $$os/$$arch"; \
			ext=""; \
			if [[ "$$os" == "windows" ]]; then ext=".exe"; fi; \
			env GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 \
				go build $(LDFLAGS) \
				-o $(DIST_DIR)/$(BINARY_NAME)-$$os-$$arch$$ext ; \
		done; \
	done

NUL = /dev/null

BUILD_DATE ?= $$(date -u +"%Y-%m-%d")
BUILD_VERSION ?= $(shell git describe --tags --abbrev=0 2>$(NUL))

LDFLAGS = -s -w
LDFLAGS += -X 'main.BuildDate=$(BUILD_DATE)'
LDFLAGS += -X "main.Version=${BUILD_VERSION}"

build:
	go build -ldflags "$(LDFLAGS)" -o bin/cypher cypher.go


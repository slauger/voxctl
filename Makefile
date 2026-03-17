BINARY     := voxctl
MODULE     := github.com/slauger/voxctl
BUILD_DIR  := bin

VERSION    ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT     ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE       ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS    := -s -w \
              -X $(MODULE)/internal/version.Version=$(VERSION) \
              -X $(MODULE)/internal/version.Commit=$(COMMIT) \
              -X $(MODULE)/internal/version.Date=$(DATE)

.PHONY: all build clean test lint vet fmt tidy snapshot install

all: build

build:
	CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o $(BUILD_DIR)/$(BINARY) ./cmd/voxctl/

install:
	CGO_ENABLED=0 go install -ldflags '$(LDFLAGS)' ./cmd/voxctl/

test:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .

lint:
	golangci-lint run

tidy:
	go mod tidy

snapshot:
	goreleaser release --snapshot --clean

clean:
	rm -rf $(BUILD_DIR) dist

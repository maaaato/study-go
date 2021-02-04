NAME := monitor
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

# 必要なツール類をセットアップする
## Setup
setup:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2
	go get github.com/client9/misspell/cmd/misspell
	go get github.com/Songmu/make2help/cmd/make2help

## Lint
lint: setup
	misspell .
	golangci-lint run ./...

## Build
build:
	GOOS=linux GOARCH=amd64 go build -o monitor
	docker build . -t monitor

## Docker run
drun:
	docker run -it --rm  --entrypoint /bin/ash monitor

## Help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup lint help
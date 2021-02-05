NAME := monitor
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'


## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# 必要なツール類をセットアップする
## Setup
.PHONY: deps
devel-deps: deps
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2
	go get github.com/client9/misspell/cmd/misspell
	go get github.com/Songmu/make2help/cmd/make2help

# テストを実行する
## Run tests
.PHONY: test
test: deps
	go test ./...


## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	misspell .
	golangci-lint run ./...

## build binaries ex. make bin/monitor
bin/%: cmd/%/main.go deps
	GOOS=linux GOARCH=amd64  go build -ldflags "$(LDFLAGS)" -o $@ $<

## build binary
.PHONY: build
build: bin/monitor

## Docker run
drun:
	docker run -it --rm  --entrypoint /bin/ash monitor

## Show help
help:
	@make2help $(MAKEFILE_LIST)
NAME := qiita
VERSION := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS := $(shell find . -type f -name '*.go')
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## Setup
setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

## Run tests
test: deps
	go test $$(glide novendor)

## Install dependencies
deps: setup
	glide install

## Update dependencies
update: setup
	glide update

## Lint
lint: setup
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

## Format source codes
fmt: setup
	goimports -x $$(glide nv -x)

## Build binaries
build: cmd/qiita/main.go deps
	go build -ldflags "$(LDFLAGS)" -o bin/qiita $<

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update test lint help
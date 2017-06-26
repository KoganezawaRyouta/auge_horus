#---------------------------------------------------------------------
# executables
#---------------------------------------------------------------------
MKDIR_P = mkdir -p

#---------------------------------------------------------------------
# rules
#---------------------------------------------------------------------
.PHONEY: all clean build run_import run_server stop_server deps help setup fmt lint setup_github_release github_release

SHELL=bash

VERSION=$(shell cat VERSION)
GOVERSION=$(shell go version)
BUILDHASH=$(shell git rev-parse --verify --short HEAD)

ROOTDIR=$(shell pwd)
BINDIR=$(ROOTDIR)/bin
DISTDIR=$(ROOTDIR)/dist
TMPDIR=$(ROOTDIR)/tmp
VENDORDIR=$(ROOTDIR)/vendor
BINARY_NAME=augehorus
BINARY=$(BINDIR)/$(BINARY_NAME)
SRC=$(shell find . -name "*.go")

LDFLAGS=-X "github.com/KoganezawaRyouta/augehorus/config.Version=$(VERSION)" -X "github.com/KoganezawaRyouta/augehorus/config.GoVersion=$(GOVERSION)" -X "github.com/KoganezawaRyouta/augehorus/config.BuildDhash=$(BUILDHASH)"
GOFLAGS=-ldflags '$(LDFLAGS)'

$(BINDIR):
	$(MKDIR_P) $@

$(TMPDIR):
	$(MKDIR_P) $@

$(VENDORDIR):
	$(MKDIR_P) $@

ENV=development

$(BINARY): $(BINDIR) $(SRC)
	@go build $(GOFLAGS) -o $@ ./cli

## build binary
build: deps $(BINARY)

## running batch
run_import:
	@$(BINARY) importer -e ${ENV}

## running batch whith help
run_help:
	@$(BINARY) help

## running batch whith version
run_version:
	@$(BINARY) version

## running api server
run_api_server:
	@$(BINARY) api_server -e ${ENV}

## running app server
run_app_server:
	@$(BINARY) app_server -e ${ENV}


## stop api server
stop_api_server:
	@kill -9 `cat ./tmp/api_server.pid`
	@echo "stop server!!, pid: " `cat ./tmp/api_server.pid`

## stop app server
stop_app_server:
	@kill -9 `cat ./tmp/app_server.pid`
	@echo "stop server!!, pid: " `cat ./tmp/app_server.pid`

## clean up tmp dir and binary
clean:
	@rm -rf $(TMPDIR)/* $(BINARY) $(VENDORDIR)/* $(DISTDIR)/*

## install dependencies
deps: setup $(VENDORDIR)
	@dep ensure -update

## format source code
fmt:
	@for pkg in $(shell go list ./... | grep -v 'vendor' | awk '{ sub("github.com/KoganezawaRyouta/augehorus", ".");print $0;}'); do \
		go fmt $$pkg; \
		done

## run golint
lint:
	@for pkg in $(shell go list ./... | grep -v 'vendor' | awk '{ sub("github.com/KoganezawaRyouta/augehorus", ".");print $0;}'); do \
		go vet $$pkg; \
		golint $$pkg; \
		done

## setup for development
setup: $(TMPDIR)
	@if [ -z `which golint 2> /dev/null` ]; then \
		go get github.com/golang/lint/golint; \
		fi
	@if [ -z `which make2help 2> /dev/null` ]; then \
		go get github.com/Songmu/make2help/cmd/make2help; \
		fi
	@if [ -z `which dep 2> /dev/null` ]; then \
		go get github.com/golang/dep/cmd/dep; \
		fi

## setup for github release
setup_github_release:
	@if [ -z `which gox 2> /dev/null` ]; then \
		go get -v github.com/mitchellh/gox; \
		fi
	@if [ -z `which ghr 2> /dev/anull` ]; then \
		go get -v github.com/tcnksm/ghr; \
		fi

## run github release please set it via `GITHUB_TOKEN` env
github_release: setup_github_release
	@gox --osarch "linux/amd64" $(GOFLAGS) --output $(TMPDIR)/$(VERSION)/"{{.OS}}_{{.Arch}}"/$(BINARY_NAME)
	@mkdir -p $(DISTDIR)/$(VERSION)
	@tar -cvzf $(DISTDIR)/$(VERSION)/$(BINARY_NAME).tar.gz $(TMPDIR)/$(VERSION)/*
	@ghr $(VERSION) $(DISTDIR)/$(VERSION)/*

## show help
help:
	@make2help $(MAKEFILE_LIST)

# ref. http://postd.cc/auto-documented-makefile/
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

COMMIT := $(shell git describe --always)
VERSION := $(shell cat version.go | perl -ne 'print "v$$1" if /Version = "(.+?)"/')
PKG := pkg
DIR := $(PKG)/$(VERSION)
TIMESTAMP := $(shell date +%s)
NAME := gosshauth

.PHONY: build
build: ## build the binary
	go build

.PHONY: clean
clean: ## clean up built binaries
	rm -fr $(PKG)

.PHONY: pkg
pkg: ## build the binaries and store into /pkg
	gox -os 'darwin linux' -arch '386 amd64' -ldflags '\
		-X main.GitCommit=$(COMMIT) \
		-X main.CompileTime=$(TIMESTAMP)' \
		-output '$(DIR)/$(NAME)_{{.OS}}_{{.Arch}}/$(NAME)'

.PHONY: release
release: pkg## release binaries at GitHub (NOTE: update version.go & the tag before this)
	bin/zip-binaries $(DIR)
	ghr -u delphinus $(VERSION) $(DIR)

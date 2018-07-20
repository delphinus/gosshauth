# ref. http://postd.cc/auto-documented-makefile/
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

COMMIT := $(shell git describe --always)
VERSION := $(shell cat version.go | perl -ne 'print "v$$1" if /Version = "(.+?)"/')
DIR := pkg/$(VERSION)
NAME := gosshauth
REPO := github.com/delphinus/$(NAME)

.PHONY: dep
dep: ## install dependencies
	go get -u github.com/blang/semver
	go get -u github.com/mitchellh/gox
	go get -u github.com/rhysd/go-github-selfupdate/selfupdate
	go get -u github.com/tcnksm/ghr
	go get -u gopkg.in/urfave/cli.v2

.PHONY: build
build: ## build the binary
	go build cmd/gosshauth.go

.PHONY: release
release: ## release binaries at GitHub (NOTE: update version.go & the tag before this)
	gox -os 'darwin linux' -arch '386 amd64' -ldflags '-X $(REPO).GitCommit=$(COMMIT)' -output '$(DIR)/$(NAME)_{{.OS}}_{{.Arch}}/$(NAME)' github.com/delphinus/gosshauth/cmd
	bin/zip-binaries $(DIR)
	ghr -u delphinus $(VERSION) $(DIR)

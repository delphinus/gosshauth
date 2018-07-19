# ref. http://postd.cc/auto-documented-makefile/
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dep
dep: ## Install dependencies
	go get gopkg.in/urfave/cli.v2
	go get github.com/rhysd/go-github-selfupdate/selfupdate
	go get github.com/blang/semver

.PHONY: build
build: ## Build the binary
	go build cmd/gosshauth.go

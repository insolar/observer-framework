include Makefile.build.mk
include Makefile.testing.mk

export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on
export GOSUMDB ?= sum.golang.org
export GOFLAGS ?= -mod=mod
export GOPROXY=https://proxy.golang.org,https://goproxy.io,https://gocenter.io,direct
export ARTIFACTS_DIR = .artifacts

#.DEFAULT_GOAL := all

.PHONY: all
all: mod clean build

.PHONY: mod
mod:
	go mod download

.PHONY: clean
clean: ## run all cleanup tasks
	go clean ./...
	rm -f $(COVERPROFILE)
	rm -rf $(BIN_DIR)

##@ Dependencies

golangci: ## install golangci-linter
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${BIN_DIR} v1.27.0

go-acc: ## install coverage tool
	go get github.com/ory/go-acc@v0.2.3

install-deps: golangci go-acc ## install necessary dependencies

.PHONY: lint
lint: ## run linter
	${BIN_DIR}/golangci-lint --color=always run ./... -v --timeout 5m

##@ Helpers

.PHONY: config
config: ## generate config
	mkdir -p $(ARTIFACTS_DIR)
	go run cmd/gen-config/gen.go

help: ## display help screen
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"}  \
		/^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

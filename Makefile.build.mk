BIN_DIR = bin
LDFLAGS ?=
BUILD_TAGS ?= "unit integration "

##@ Building
.PHONY: build
build: ## build all applications

.PHONY: generate
generate: ## generate mocks
	GOFLAGS="" go generate ./...

.PHONY: generate-protobuf
generate-protobuf: ## generate protobuf structs
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi

.PHONY: install-dev
install-dev:
	go get github.com/gogo/protobuf

.PHONY: compile-tests
compile-tests: ## compile all tests
	go list -tags=$(BUILD_TAGS) ./... | xargs -n 1 go test -tags=$(BUILD_TAGS) -c

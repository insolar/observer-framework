COVERPROFILE ?= coverage.out

TEST_COUNT ?= 1
TEST_ARGS ?=

##@ Testing

.PHONY: unit
unit:  ## run unit tests
	go test -v ./... -tags unit -count $(TEST_COUNT) -race $(TEST_ARGS)

.PHONY: test
test: unit integration  ## run all tests

.PHONY: integration
integration: ## run integrations tests with race
	go test -v ./... -tags integration -count $(TEST_COUNT) -race $(TEST_ARGS)

.PHONY: test-with-coverage
test-with-coverage: ## run tests with coverage mode
	go-acc --covermode=count --output=coverage.tmp.out ./... -- -tags "unit integration heavy_mock_integration" -count=1
	cat coverage.tmp.out | grep -v _mock.go > ${COVERPROFILE}
	go tool cover -html=${COVERPROFILE} -o coverage.html


##@ Benchmarks

.PHONY: bench
bench: ## run benchmarks
	go test -v ./... -tags bench -bench=. -benchmem -benchtime=1000x

.PHONY: bench-compare
bench-compare: ## run benchmarks compare for last two commits
	cob -bench-cmd make -bench-args bench -threshold 0.7

.PHONY: bench-integration
bench-integration: ## run integration benchmarks
	go test -v ./... -tags bench_integration -bench=. -benchmem -benchtime=1x

.PHONY: bench-compare-integration
bench-compare-integration: ## run integration benchmarks
	cob -bench-cmd make -bench-args bench-integration -threshold 0.7

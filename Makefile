NAME ?= gommon
VERSION ?= 0.14.0

.PHONY: version tests

version:
	@echo $(VERSION)

tests:
	@echo "Running unit tests..."
	@go test ./... -coverprofile cover.out
	@go tool cover -html=cover.out -o coverage.html
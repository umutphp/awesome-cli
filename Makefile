SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=
OS=$(shell uname -s)

export PATH := ./bin:$(PATH)

# Install all the build and lint dependencies
setup:
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
ifeq ($(OS), Darwin)
    brew install dep
else
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
    dep ensure -vendor-only
.PHONY: setup

test:
    go test $(TEST_OPTIONS) -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m

cover: test
    go tool cover -html=coverage.out

fmt:
    find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint:
    golangci-lint run --enable-all ./...

ci: lint test

.DEFAULT_GOAL := build
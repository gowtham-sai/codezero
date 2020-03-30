.Phony: all

SHELL := /bin/bash # Use bash syntax
default: check-quality test build

ALL_PACKAGES=$(shell go list ./... | grep -v /vendor)
WORKDIR=$(shell echo "${PWD}")
APP_EXECUTABLE="out/codezero"

setup: --setup-git-hooks
	go get -u golang.org/x/lint/golint
	go get gotest.tools/gotestsum

--setup-git-hooks:
	@if [ -d "$$PWD/.git/hooks" ]; then \
	  if [ -L "$$PWD/.git/hooks" ]; then \
		rm "$$PWD/.git/hooks"; \
	  else \
		rm -rf "$$PWD/.git/hooks"; \
	  fi \
	fi
	ln -s "$$PWD/.githooks" "$$PWD/.git/hooks"


check-quality: --lint --vet

--lint:
	@echo "Running Lint ...."
	@if [[ `golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } | wc -l | tr -d ' '` -ne 0 ]]; then \
          golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }; \
          exit 2; \
    fi;

--vet:
	@echo "Running Vet checks ...."
	@go vet ./...

test:
	GO111MODULE=on go clean -testcache ./... && go test ./...

build:
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE) .

test-coverage:
	GO111MODULE=on go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

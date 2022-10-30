HOSTNAME=liquidcollective.io
NAMESPACE=dev
NAME=fireblocks
BINARY=$(BUILD_FOLDER)/terraform-provider-${NAME}
VERSION=0.1.31

OS_ARCH=darwin_arm64

GOPATH ?= $(shell go env GOPATH)

# List of effective go files
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*" -not -path "./tests/*" | egrep -v "^\./\.go" | grep -v _test.go)

# List of packages except testsutils
PACKAGES ?= $(shell go list ./... | grep -v "mock" )

# Build folder
BUILD_FOLDER = build

# Test coverage variables
COVERAGE_BUILD_FOLDER = $(BUILD_FOLDER)/coverage

UNIT_COVERAGE_OUT = $(COVERAGE_BUILD_FOLDER)/ut_cov.out
UNIT_COVERAGE_HTML =$(COVERAGE_BUILD_FOLDER)/ut_index.html

INTEGRATION_COVERAGE_OUT = $(COVERAGE_BUILD_FOLDER)/it_cov.out
INTEGRATION_COVERAGE_HTML =$(COVERAGE_BUILD_FOLDER)/it_index.html

# Test lint variables
GOLANGCI_VERSION = v1.50.1

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OPEN = xdg-open
endif
ifeq ($(UNAME_S),Darwin)
	OPEN = open
endif

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build/coverage:
	@mkdir -p build/coverage

unit-test: build/coverage
	@TF_ACC=1 go test -covermode=count -coverprofile $(UNIT_COVERAGE_OUT) $(PACKAGES)

unit-test-cov: unit-test
	@go tool cover -html=$(UNIT_COVERAGE_OUT) -o $(UNIT_COVERAGE_HTML)

integration-test: build/coverage
	@go test -covermode=count -coverprofile $(INTEGRATION_COVERAGE_OUT) -v --tags integration ${PACKAGES}

integration-test-cov: integration-test
	@go tool cover -html=$(INTEGRATION_COVERAGE_OUT) -o $(INTEGRATION_COVERAGE_HTML)
	
test-race:
	@TF_ACC=1 go test -race $(PACKAGES)

fix-lint: ## Run linter to fix issues
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:$(GOLANGCI_VERSION) golangci-lint run --fix

# @misspell -error $(GOFILES)
test-lint: ## Check linting
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:$(GOLANGCI_VERSION) golangci-lint run -v --out-format=github-actions -c .golangci.yml --allow-parallel-runners

mockgen-install:
	@type mockgen >/dev/null 2>&1 || {   \
		echo "Installing mockgen..."; \
		go install github.com/golang/mock/mockgen@v1.6.0;  \
	}

mockgen: mockgen-install
	$(GOPATH)/bin/mockgen -source pkg/fireblocks/client/client.go -destination pkg/fireblocks/client/mock/client.go -package mock Client

generate:
	$(GOPATH)/bin/tfplugindocs generate --provider-name "terraform-provider-fireblocks" --rendered-provider-name "Fireblocks" --tf-version "v1.3.1"

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

go-build:
	go build -o ${BINARY}

install: go-build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

init:
	cd examples && terraform init -upgrade

plan:
	cd examples && terraform plan

apply:
	cd examples && terraform apply -auto-approve

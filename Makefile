.PHONY: all default build clean vendor watch debug start test coverage test-summary lint lint-dockerfile lint-go lint-yaml format examine fix docker-build docker-release help

# Shell to use for running scripts
SHELL := $(shell which bash)
GOCMD := go
GOTEST := $(GOCMD) test
GOVET := $(GOCMD) vet
GOFMT := $(GOCMD) fmt
GOFIX := $(GOCMD) fix
BINARY_NAME := rfdez/my-game
VERSION ?= 0.1.0
DOCKER_REGISTRY ?= ghcr.io/#if set it should finished by /
EXPORT_RESULT ?= false# for CI please set EXPORT_RESULT to true

# Test if the dependencies we need to run this Makefile are installed
DOCKER := $(shell command -v docker)
DOCKER_COMPOSE := $(shell command -v docker-compose)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: help

default:
ifndef DOCKER
	@echo "Docker is not available. Please install docker"
	@exit 1
endif
ifndef DOCKER_COMPOSE
	@echo "docker-compose is not available. Please install docker-compose"
	@exit 1
endif

## Build:
build: ## Build your project and put the output binary in out/bin/
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "mkdir -p out/bin \
		&& GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(BINARY_NAME) cmd/api/main.go"

clean: ## Remove build related file
	@rm -fr ./bin
	@rm -fr ./out
	@rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml go_test.xml

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "$(GOCMD) mod vendor"

watch: ## Run the code with cosmtrek/air to have automatic reload on changes
	@docker-compose up dev-app && docker-compose down --remove-orphans --rmi local

debug: ## Run the code with dlv to have a debug session
	@docker-compose up --build debug-app && docker-compose down --remove-orphans --rmi local

start: ## Run the code with docker-compose
	@docker-compose up --build app && docker-compose down --remove-orphans --rmi local

## Test:
test: ## Run the tests
ifeq ($(EXPORT_RESULT), true)
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "GO111MODULE=off go get -u github.com/jstemmer/go-junit-report \
		&& $(GOTEST) -v -race ./... | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml"
else
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "$(GOTEST) -v -race ./..."
endif

coverage: ## Run the tests and generate coverage report
ifeq ($(EXPORT_RESULT), true)
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "GO111MODULE=off go get -u github.com/AlekSi/gocov-xml \
		&& GO111MODULE=off go get -u github.com/axw/gocov/gocov \
		&& $(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./... \
		&& go tool cover -func profile.cov \
		&& gocov convert profile.cov | gocov-xml - > coverage.xml"
else
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./... \
		&& go tool cover -func profile.cov"
endif

test-summary: ## Run the tests with the summary
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "GO111MODULE=off go get -u gotest.tools/gotestsum \
		&& gotestsum --junitfile go_test.xml --format testname"

## Lint:
lint: lint-go lint-dockerfile lint-yaml ## Run all available linters

lint-dockerfile: ## Run the linter on the dockerfile
# If dockerfile is present we lint it.
ifeq ($(shell test -e ./Dockerfile && echo -n yes),yes)
	$(eval CONFIG_OPTION = $(shell [ -e $(shell pwd)/.hadolint.yaml ] && echo "-v $(shell pwd)/.hadolint.yaml:/root/.config/hadolint.yaml" || echo "" ))
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--format checkstyle" || echo "" ))
	$(eval OUTPUT_FILE = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "| tee /dev/tty > checkstyle-report.xml" || echo "" ))
	@docker run --rm -i $(CONFIG_OPTION) hadolint/hadolint:latest hadolint $(OUTPUT_OPTIONS) - < ./Dockerfile $(OUTPUT_FILE)
endif

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	@docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)

lint-yaml: ## Use yamllint on the yaml file of your projects
ifeq ($(EXPORT_RESULT), true)
	@GO111MODULE=off go get -u github.com/thomaspoignant/yamllint-checkstyle
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | yamllint-checkstyle > yamllint-checkstyle.xml)
endif
	@docker run --rm -it -v $(shell pwd):/data cytopia/yamllint:latest -f parsable $(shell git ls-files '*.yml' '*.yaml') $(OUTPUT_OPTIONS)

format: ## Fromat the code with gofmt
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "$(GOFMT) ./..."

examine: ## Examine the code with govet
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "$(GOVET) ./..."

fix: ## Fix the code with gofix
	@docker run --rm -v $(PWD):/app -w /app golang:1.18 sh -c "$() ./..."

## Docker:
docker-build: ## Use the dockerfile to build the container
	@docker build --rm --tag $(BINARY_NAME) .

docker-release: ## Release the container with tag latest and version
	@docker tag $(BINARY_NAME) $(DOCKER_REGISTRY)$(BINARY_NAME):latest
	@docker tag $(BINARY_NAME) $(DOCKER_REGISTRY)$(BINARY_NAME):$(VERSION)
	# Push the docker images
	@docker push $(DOCKER_REGISTRY)$(BINARY_NAME):latest
	@docker push $(DOCKER_REGISTRY)$(BINARY_NAME):$(VERSION)

docker-clean: ## Clean the docker resources
	@docker-compose down --remove-orphans --rmi local --volumes

docker-dbclient: ## Build the dbclient container
	@docker-compose up adminer && docker-compose down --remove-orphans --rmi local

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

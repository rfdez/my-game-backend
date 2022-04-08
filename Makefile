.PHONY = deps build

# Shell to use for running scripts
SHELL := $(shell which bash)

# Test if the dependencies we need to run this Makefile are installed
DOCKER := $(shell command -v docker)
DOCKER_COMPOSE := $(shell command -v docker-compose)
GO := $(shell command -v go)
deps:
ifndef DOCKER
	@echo "Docker is not available. Please install docker"
	@exit 1
endif
ifndef DOCKER_COMPOSE
	@echo "docker-compose is not available. Please install docker-compose"
	@exit 1
endif
ifndef GO
	@echo "Go is not available. Please install Go"
	@exit 1
endif

build: deps
	@docker-compose up --build

test: deps
	@go test -v ./...

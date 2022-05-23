.PHONY: version git-version init run

_GIT_LAST_COMMIT_TIME=$(shell TZ=UTC git log --pretty=format:'%cd' -1 --date=format-local:'%Y%m%d-%H%M%S')
_GIT_LAST_COMMIT_HASH=$(shell git rev-parse --short HEAD)
_GIT_VERSION=$(_GIT_LAST_COMMIT_TIME).$(_GIT_LAST_COMMIT_HASH)

_VERSION=$(shell cat Version)

DOCKER_IMAGE_PREFIX=
DOCKER_IMAGE_NAME=$(DOCKER_IMAGE_PREFIX)golang-project-demo:$(_VERSION)

GOCMD=go
GOOS=linux
GOARCH=amd64
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GOLDFLAGS += -X main.Version=$(_VERSION)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

version:
	@echo ${_VERSION}

git-version:
	@echo ${_GIT_VERSION}

init:
	@mkdir -p bin

build: image-build

run: container-run

stop: container-stop

# Local
local-run:
	@$(GORUN) $(GOFLAGS) -v ./cmd

local-build: init
	@$(GOBUILD) $(GOFLAGS) -o bin/demo -v ./cmd

# Docker
image-build: local-build
	docker build -t $(DOCKER_IMAGE_NAME) .

container-run: build
	DOCKER_IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker-compose up -d

container-stop:
	DOCKER_IMAGE_NAME=$(DOCKER_IMAGE_NAME) docker-compose down -v

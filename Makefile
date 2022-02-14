REVISION := $(shell git rev-parse --short=8 HEAD || echo unknown)
VERSION ?= $(shell git describe --tags --always)
BUILT := $(shell date +%Y-%m-%dT%H:%M:%S)
BRANCH := $(shell git show-ref | grep "$(REVISION)" | grep -v HEAD | awk '{print $$2}' | sed 's|refs/remotes/origin/||' | sed 's|refs/heads/||' | sort | head -n 1)
MODULE_NAME := $(shell cat go.mod | head -n 1 | cut -d " " -f2- | cut -d "/" -f2-)

GO_LDFLAGS ?= -X main.version=$(VERSION) \
			-X main.branch=$(BRANCH) \
			-X main.revision=$(REVISION) \
			-X main.built=$(BUILT) \
			-s \
			-w

BUILD_DIR=build
DIST_DIR=dist

build: clean
	go build -a -ldflags "$(GO_LDFLAGS)" -o="$(BUILD_DIR)/todoapp" ./cmd/todo-app

build-cli: clean dist-cli-darwin dist-cli-linux dist-cli-windows

clean: 
	@rm -rf ./${BUILD_DIR} ./${DIST_DIR}

protos:
	@echo $(shell pwd)
	docker build -f Docker.buf -t buf .
	docker run --rm -v $(shell pwd):/workspace buf generate --template buf.gen.go.yaml \
	 --path protobuf/protos/api \

ONESHELL:
.SILENT:
dist-cli-prepare:
	mkdir -p ${DIST_DIR}

.ONESHELL:
.SILENT:
dist-cli-darwin: dist-cli-prepare
	GOOS=darwin GOARCH=amd64 go build -a -ldflags "$(GO_LDFLAGS)" -o="$(DIST_DIR)/darwin/todoapp" ./cmd/todo-app
	
.ONESHELL:
.SILENT:
dist-cli-linux: dist-cli-prepare
	GOOS=linux GOARCH=amd64 go build -a -ldflags "$(GO_LDFLAGS)" -o="$(DIST_DIR)/linux/todoapp" ./cmd/todo-app

.ONESHELL:
.SILENT:
dist-cli-windows: dist-cli-prepare
	GOOS=windows GOARCH=amd64 go build -a -ldflags "$(GO_LDFLAGS)" -o="$(DIST_DIR)/windows/todoapp.exe" ./cmd/todo-app
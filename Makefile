.DEFAULT_GOAL := help

DOCKER_RUN := docker run --rm -v "$(CURDIR):/work" -w /work

GO_IMAGE := golang:1.26.9-trixie@sha256:f06b7122f888680048afb247c4e5f49c0069eb6c4c1ccae0f3f4834527a8d22c
# Version written into the binary: the nearest tag, or the commit when there is none.
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
# Named volumes keep the Go module and build caches between runs.
GO_RUN := $(DOCKER_RUN) -v $(notdir $(CURDIR))-go-mod:/go/pkg/mod -v $(notdir $(CURDIR))-go-build:/root/.cache/go-build

.PHONY: help test build

help: ## List available targets
	@grep -E '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  %-18s %s\n", $$1, $$2}'

test: ## Run the Go tests with the race detector
	$(GO_RUN) $(GO_IMAGE) go test -race ./...

# go build stamps the git commit into the binary. In CI the checkout belongs to
# another user than the container, so git needs /work marked as safe.
build: ## Build simurg for Linux and macOS on amd64 and arm64 into dist/
	$(GO_RUN) -e VERSION=$(VERSION) \
		-e GIT_CONFIG_COUNT=1 -e GIT_CONFIG_KEY_0=safe.directory -e GIT_CONFIG_VALUE_0=/work \
		$(GO_IMAGE) sh scripts/ci/build.sh

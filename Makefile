SHELL = /bin/bash
ifeq ($(shell uname -s),Windows)
	.SHELLFLAGS = /o pipefile /c
else
	.SHELLFLAGS = -o pipefail -c
endif
# SRC_ROOT is the top of the source tree.
SRC_ROOT := $(shell git rev-parse --show-toplevel)
GOCMD?= go
GOTEST=$(GOCMD) test
GOTEST_OPT?= -race -timeout 300s -parallel 4 --tags=$(GO_BUILD_TAGS)
TOOLS_BIN_DIR    := $(SRC_ROOT)/.tools
TOOLS_MOD_DIR    := $(SRC_ROOT)/internal/tools
TOOLS_MOD_REGEX  := "\s+_\s+\".*\""
TOOLS_PKG_NAMES  := $(shell grep -E $(TOOLS_MOD_REGEX) < $(TOOLS_MOD_DIR)/tools.go | tr -d " _\"")
TOOLS_BIN_DIR    := $(SRC_ROOT)/.tools
TOOLS_BIN_NAMES  := $(addprefix $(TOOLS_BIN_DIR)/, $(notdir $(TOOLS_PKG_NAMES)))
DIST_DIR		 := $(SRC_ROOT)/dist

.PHONY: all
all: install-tools gotidy golint govulncheck gotest

.PHONY: install-tools
install-tools: $(TOOLS_BIN_NAMES)
$(TOOLS_BIN_DIR):
	mkdir -p $@

$(TOOLS_BIN_NAMES): $(TOOLS_BIN_DIR) $(TOOLS_MOD_DIR)/go.mod
	cd $(TOOLS_MOD_DIR) && $(GOCMD) build -o $@ -trimpath $(filter %/$(notdir $@),$(TOOLS_PKG_NAMES))

LINT                := $(TOOLS_BIN_DIR)/golangci-lint
GOTESTSUM           := $(TOOLS_BIN_DIR)/gotestsum
GOVULNCHECK         := $(TOOLS_BIN_DIR)/govulncheck
GORELEASER			:= $(TOOLS_BIN_DIR)/goreleaser
GOIMPORTS           := $(TOOLS_BIN_DIR)/goimports

.PHONY: golint
golint:
	$(LINT) run --allow-parallel-runners --build-tags integration --path-prefix $(shell basename "$(CURDIR)")


.PHONY: gotest
gotest:
	  $(GOTEST) $(GOTEST_OPT) ./...

.PHONY: gotidy
gotidy:
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.21

.PHONY: govulncheck
govulncheck: $(GOVULNCHECK)
	$(GOVULNCHECK) ./...

.PHONY: gomoddownload
moddownload:
	$(GOCMD) mod download

.PHONY: build
build:
	$(GORELEASER) build  --clean --skip-validate
.PHONY: cleanup
cleanup:
	if [ -d  $(TOOLS_BIN_DIR) ]; then rm -r $(TOOLS_BIN_DIR); fi
	if [ -d  $(DIST_DIR) ]; then rm -r $(DIST_DIR); fi
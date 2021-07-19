# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.4.3. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for bingo variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(BINGO)
#	@echo "Running bingo"
#	@$(BINGO) <flags/args..>
#
BINGO := $(GOBIN)/bingo-v0.4.3
$(BINGO): $(BINGO_DIR)/bingo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/bingo-v0.4.3"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=bingo.mod -o=$(GOBIN)/bingo-v0.4.3 "github.com/bwplotka/bingo"

NATS_SERVER := $(GOBIN)/nats-server-v2.3.2
$(NATS_SERVER): $(BINGO_DIR)/nats-server.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/nats-server-v2.3.2"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=nats-server.mod -o=$(GOBIN)/nats-server-v2.3.2 "github.com/nats-io/nats-server/v2"

NATS := $(GOBIN)/nats-v0.0.0-20201207130909-c15ad280ec4b
$(NATS): $(BINGO_DIR)/nats.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/nats-v0.0.0-20201207130909-c15ad280ec4b"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=nats.mod -o=$(GOBIN)/nats-v0.0.0-20201207130909-c15ad280ec4b "github.com/nats-io/natscli"

NSC := $(GOBIN)/nsc-v0.0.0-20210521221149-a973cba21762
$(NSC): $(BINGO_DIR)/nsc.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/nsc-v0.0.0-20210521221149-a973cba21762"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=nsc.mod -o=$(GOBIN)/nsc-v0.0.0-20210521221149-a973cba21762 "github.com/nats-io/nsc"


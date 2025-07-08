VERSION = 0.1.6

DVCS_HOST = github.com
ORG = geomyidia
PROJ = zylog
FQ_PROJ = $(DVCS_HOST)/$(ORG)/$(PROJ)

GO ?= go
DEFAULT_GOPATH=$(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
export PATH:=$(PATH):$(DEFAULT_GOBIN)

GOLANGCI_LINT=$(DEFAULT_GOBIN)/golangci-lint

LD_VERSION = -X $(FQ_PROJ)/logger.Version=$(VERSION)
LD_BUILDDATE = -X $(FQ_PROJ)/logger.BuildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_GITCOMMIT = -X $(FQ_PROJ)/logger.GitCommit=$(shell git rev-parse --short HEAD)
LD_GITBRANCH = -X $(FQ_PROJ)/logger.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)
LD_GITSUMMARY = -X $(FQ_PROJ)/logger.GitSummary=$(shell git describe --tags --dirty --always)

LDFLAGS = -w -s $(LD_VERSION) $(LD_BUILDDATE) $(LD_GITBRANCH) $(LD_GITSUMMARY) $(LD_GITCOMMIT)

default: lint build

default-gopath:
	@echo $(DEFAULT_GOPATH)

build:
	GO111MODULE=on $(GO) build \
		-ldflags "$(LDFLAGS)" \
		-o ./bin/zylog-demo \
		github.com/geomyidia/zylog/cmd/zylog-demo

modules-init:
	GO111MODULE=on $(GO) mod init github.com/geomyidia/zylog

modules-update:
	GO111MODULE=on $(GO) get -u

install-goimports:
	GO111MODULE=on $(GO) get golang.org/x/tools/cmd/goimports

$(GOLANGCI_LINT):
	@echo "Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v2.2.1

lint: $(GOLANGCI_LINT)
	@echo '>> Linting source code'
	@echo "($$($(GOLANGCI_LINT) --version))"
	@$(GOLANGCI_LINT) run \
	--enable=errcheck \
	--enable=dupl \
	--enable=unparam \
	--enable=wastedassign \
	--enable=ineffassign \
	--enable=revive \
	--enable=gocritic \
	--enable=misspell \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	--enable=govet \
	--show-stats \
	./...

rm-lint:
	@echo "Removing golangci-lint ..."
	rm -f $(GOLANGCI_LINT)

goimports:
	GO111MODULE=on goimports -v -w ./

clean:
	rm -rf bin/*

clean-all:
	$(GO) clean --modcache

show-ldflags:
	@echo $(LDFLAGS)

upgrade-deps:
	@echo ">> Upgrading Go module dependencies ..."
	@$(GO) get -u ./...
	@$(GO) mod tidy

TEST_RUNNER = $(DEFAULT_GOBIN)/gotestsum
$(TEST_RUNNER):
	@echo ">> Couldn't find $(TEST_RUNNER); installing ..."
	GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get gotest.tools/gotestsum && \
	$(GO) install gotest.tools/gotestsum

test: $(TEST_RUNNER)
	@echo '>> Running all tests'
	@GO111MODULE=on $(TEST_RUNNER) --format testname -- ./...

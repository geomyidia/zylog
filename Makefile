VERSION = 0.1.5

DVCS_HOST = github.com
ORG = geomyidia
PROJ = zylog
FQ_PROJ = $(DVCS_HOST)/$(ORG)/$(PROJ)

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
	GO111MODULE=on go build \
		-ldflags "$(LDFLAGS)" \
		-o ./bin/zylog-demo \
		github.com/geomyidia/zylog/cmd/zylog-demo

modules-init:
	GO111MODULE=on go mod init github.com/geomyidia/zylog

modules-update:
	GO111MODULE=on go get -u

install-goimports:
	GO111MODULE=on go get golang.org/x/tools/cmd/goimports

$(GOLANGCI_LINT):
	@echo "Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.15.0

lint: $(GOLANGCI_LINT)
	GO111MODULE=on golangci-lint \
	--enable=gofmt \
	--enable=golint \
	--enable=gocritic \
	--enable=misspell \
	--enable=nakedret \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	run ./...

goimports:
	GO111MODULE=on goimports -v -w ./

clean:
	rm -rf bin/*

clean-all:
	go clean --modcache

show-ldflags:
	@echo $(LDFLAGS)

upgrade-deps:
	@echo ">> Upgrading Go module dependencies ..."
	@go get -u ./...
	@go mod tidy

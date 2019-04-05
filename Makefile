VERSION = 0.1.3
BUILD_FLAGS=$(shell govvv -flags -pkg github.com/zylisp/zylog/logger -version $(VERSION))
GOLANGCI_LINT=$(shell which golangci-lint)
DEFAULT_GOPATH=$(shell tr ':' '\n' <<< "$$GOPATH"|sed '/^$$/d'|head -1)
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin

default: lint

default-gopath:
	@echo $(DEFAULT_GOPATH)

build-deps:
	go get github.com/ahmetb/govvv

build:
# 	@go build -ldflags="$(BUILD_FLAGS)" github.com/zylisp/zylog/logger
	@GO111MODULE=on go build \
		-o ./bin/zylog-demo \
		-ldflags="$(BUILD_FLAGS)" \
		github.com/zylisp/zylog/cmd/zylog-demo


modules-init:
	GO111MODULE=on go mod init github.com/zylisp/zylog

modules-update:
	GO111MODULE=on go get -u

install-goimports:
	GO111MODULE=on go get golang.org/x/tools/cmd/goimports

$(GOLANGCI_LINT):
	@echo "Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) latest

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

VERSION = 0.2.0
BUILD_FLAGS=$(shell govvv -flags -pkg github.com/zylisp/zylog/logger -version $(VERSION))

default: lint

build-deps:
	go get github.com/ahmetb/govvv

build:
# 	@go build -ldflags="$(BUILD_FLAGS)" github.com/zylisp/zylog/logger
	@go build \
		-o ./bin/demo \
		-ldflags="$(BUILD_FLAGS)" \
		github.com/zylisp/zylog/cmd/demo


modules-init:
	GO111MODULE=on go mod init github.com/zylisp/zylog

modules-update:
	GO111MODULE=on go get -u

install-goimports:
	GO111MODULE=on go get golang.org/x/tools/cmd/goimports

install-golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b `go env GOPATH`/bin latest

lint:
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

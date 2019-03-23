default: lint

modules-init:
	GO111MODULE=on go mod init github.com/MediaMath/identity-1p3p-service

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

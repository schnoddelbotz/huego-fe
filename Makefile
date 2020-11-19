
BINARY := huego-fe
VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null || echo v0)
# ^ bring back: --dirty ... after throwing out assets.go after https://github.com/golang/go/issues/41191
LDFLAGS := -w -s -X github.com/schnoddelbotz/huego-fe/cmd.Version=$(VERSION)
WIN_LDFLAGS := -H=windowsgui
ASSETS := $(wildcard assets/*)
SRC_DEPENDS := main.go */*.go go.* web/assets.go
OS := $(shell uname -s)

build: web/assets.go $(BINARY)

$(BINARY): $(BINARY)_$(OS)

$(BINARY)_Linux: $(SRC_DEPENDS)
	# building for Linux
	GOOS=linux GOARCH=amd64 go build -ldflags='$(LDFLAGS)'

$(BINARY)_Windows: $(SRC_DEPENDS)
	# building for Windows; Using "-H=windowsgui" will disable console window, but also break CLI usage.
	# Dunno how to fix. Use `WINGUI="" make huego-fe_Windows` if CLI version is desired.
	GOOS=windows GOARCH=amd64 go build -ldflags='$(LDFLAGS) $(WIN_LDFLAGS)'

$(BINARY)_Darwin: $(SRC_DEPENDS)
	# building for the OS fka MacOS X
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags='$(LDFLAGS)'

web/assets.go: $(ASSETS)
	# building web/assets.go to embed web assets into huego-fe binary
	test -n "$(shell which esc)" || go get -v -u github.com/mjibson/esc
	go generate

fmt:
	go fmt ./...

clean:
	rm -f $(BINARY) $(BINARY).exe web/assets.go

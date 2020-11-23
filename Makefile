
BINARY := huego-fe
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo v0)
LDFLAGS := -w -s -X github.com/schnoddelbotz/huego-fe/cmd.Version=$(VERSION)
WIN_LDFLAGS := -H=windowsgui
ASSETS := $(wildcard assets/*)
SRC_DEPENDS := main.go */*.go go.* web/assets.go
COVERAGE_PROFILE := profile.cov

OS := $(shell uname -s)
# Linux: X11 or Wayland ?
LINUX_DISPLAY := X11
# Windows: CLI or GUI? Set empty for CLI.
WIN_UI := -H=windowsgui

all: test $(BINARY)

$(BINARY): $(SRC_DEPENDS)
	# building $(BINARY) for OS $(OS)
	$(MAKE) $(BINARY)_$(OS)

$(BINARY)_Linux: $(BINARY)_Linux_$(LINUX_DISPLAY)

$(BINARY)_Linux_X11:
	# building for Linux/X11
	GOOS=linux GOARCH=amd64 go build -tags nowayland -ldflags='$(LDFLAGS)'

$(BINARY)_Linux_Wayland:
	# building for Linux/Wayland
	GOOS=linux GOARCH=amd64 go build -tags nox11 -ldflags='$(LDFLAGS)'

$(BINARY)_Darwin:
	# building for the OS fka MacOS X
	GOOS=darwin GOARCH=amd64 go build -ldflags='$(LDFLAGS)'

$(BINARY).exe:
	# building for Windows; Using "-H=windowsgui" will disable console window, but also break CLI usage.
	# Dunno how to fix. Use `make huego-fe.exe WINGUI=''` if CLI version is desired.
	GOOS=windows GOARCH=amd64 go build -ldflags='$(LDFLAGS) $(WINDOWS_UI)'

web/assets.go: $(ASSETS)
	# building web/assets.go to embed web assets into huego-fe binary
	go generate

fmt:
	go fmt ./...

lint:
	golint ./...

ineffassign:
	ineffassign .

vet:
	go vet .

test: web/assets.go
	go test -race -covermode atomic -coverprofile=$(COVERAGE_PROFILE) ./...

test_all: lint ineffassign vet test

clean:
	rm -f $(BINARY) $(BINARY).exe $(COVERAGE_PROFILE)

git-setup:
	/bin/echo -e '#!/bin/sh\nexec make fmt' > .git/hooks/pre-commit
	/bin/echo -e '#!/bin/sh\nexec make test_all' > .git/hooks/pre-push
	chmod +x .git/hooks/pre-commit .git/hooks/pre-push

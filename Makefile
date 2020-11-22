
BINARY := huego-fe
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo v0)
LDFLAGS := -w -s -X github.com/schnoddelbotz/huego-fe/cmd.Version=$(VERSION)
WIN_LDFLAGS := -H=windowsgui
ASSETS := $(wildcard assets/*)
SRC_DEPENDS := main.go */*.go go.* web/assets.go

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
	test -n "$(shell which esc)" || go get -v -u github.com/mjibson/esc
	go generate

fmt:
	go fmt ./...

test: web/assets.go
	go test ./...

clean:
	rm -f $(BINARY) $(BINARY).exe


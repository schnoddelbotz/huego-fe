
BINARY := huego-fe
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo v0)
LDFLAGS := -X github.com/schnoddelbotz/huego-fe/cmd.Version=$(VERSION)
ASSETS := $(wildcard assets/*)

all: $(BINARY)


$(BINARY): main.go */*.go go.* web/assets.go
	# CGO_ENABLED=0  <- OK on mac, fail on Linux atm ... tbd.
	go build -ldflags='-w -s $(LDFLAGS)'

web/assets.go: $(ASSETS)
	test -n "$(shell which esc)" || go get -v -u github.com/mjibson/esc
	go generate

clean:
	rm -f $(BINARY) web/assets.go

fmt:
	go fmt ./...

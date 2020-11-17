
BINARY := huego-fe
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo v0)
LDFLAGS := -X github.com/schnoddelbotz/huego-fe/cmd.Version=$(VERSION)

all: $(BINARY)


$(BINARY): main.go */*.go go.*
	CGO_ENABLED=0 go build -ldflags='-w -s $(LDFLAGS)'

clean:
	rm -f $(BINARY)

fmt:
	go fmt ./...

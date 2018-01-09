VERSION := $(shell git describe --abbrev=0 --tags)

all: release

release:
	go build -ldflags "-X main.version=$(VERSION)"

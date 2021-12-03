BASE ?= b
EXTENSION ?= e

build-base:
	go build -o $(BASE) base/main.go

build-extension:
	go build -o $(EXTENSION) extension/main.go

build: build-base build-extension

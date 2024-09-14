CMD := fasterx
PROG ?= figma


.PHONY: build
build:
	go build -o bin/$(CMD)

.PHONY: run
run:
	bin/${CMD} $(PROG)

.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/fasterx-windows-amd64.exe

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/fasterx-linux-amd64

.PHONY: build-mac-i
build-mac-i:
	GOOS=darwin GOARCH=amd64 go build -o bin/fasterx-mac-amd64

.PHONY: build-mac-m
build-mac-m:
	GOOS=darwin GOARCH=arm64 go build -o bin/fasterx-mac-arm64

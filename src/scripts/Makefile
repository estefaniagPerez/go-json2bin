.DEFAULT_GOAL := build
fmt:
	go fmt ../...
.PHONY:fmt

lint: fmt
	~/go/bin/golint ../...
.PHONY:lint

vet: fmt
	go vet ../...
	~/go/bin/shadow ../...
.PHONY:vet

build: vet
	go build ../cmd/json2bin-parser/
.PHONY:build
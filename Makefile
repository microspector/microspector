export PWD=$(shell pwd)
export GO111MODULE=on

build:
	go build -o $(PWD)/bin/microspector $(PWD)/cmd/main.go

parser: deps
	$(GOPATH)/bin/pigeon ./pkg/parser/msf.peg  > ./pkg/parser/msf.go

deps:
	go get

test:
	go test

.PHONY: build parser deps test
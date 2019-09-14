export PWD=$(shell pwd)
export GO111MODULE=on

build: parser
	go build -o $(PWD)/bin/microspector $(PWD)/cmd

parser: deps
	$(GOPATH)/bin/pigeon ./pkg/parser/msf.peg  > ./pkg/parser/msf.go

yacc:
	goyacc -o ./pkg/parser/parser.go ./pkg/parser/parser.y && go run ./cmd

deps:
	go get

test:
	go test $(PWD)/pkg/parser

.PHONY: build parser deps test
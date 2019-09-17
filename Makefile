export PWD=$(shell pwd)
export GO111MODULE=on

build: parser
	go build -o $(PWD)/bin/microspector $(PWD)/cmd

parser: deps
	$(GOPATH)/bin/pigeon ./pkg/parser/msf.peg  > ./pkg/parser/msf.go

run: yacc
	go run ./cmd --file="tasks/main.msf"

test: yacc
	go test $(PWD)/pkg/parser

yacc:
	goyacc -o ./pkg/parser/parser.go ./pkg/parser/parser.y
	gofmt -w ./pkg/parser/parser.go

deps:
	go get -v all

.PHONY: build parser deps test
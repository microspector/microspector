export PWD=$(shell pwd)
export GO111MODULE=on

buildwin: yacc
	GOOS=windows go build -o $(PWD)/bin/microspector.exe $(PWD)/cmd

build: yacc
	go build -o $(PWD)/bin/microspector $(PWD)/cmd

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
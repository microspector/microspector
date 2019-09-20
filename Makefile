PWD := $(shell pwd)
export GO111MODULE=on
VERSION := $(shell git describe --always --long --dirty)

buildwin: yacc
	GOOS=windows go build -i -v -o ${PWD}/bin/microspector.exe -ldflags="-X main.version=${VERSION}" ${PWD}/cmd

build: yacc
	go build -i -v -o $(PWD)/bin/microspector -ldflags="-X main.version=${VERSION}" ${PWD}/cmd

run: yacc
	go run ./cmd --folder="tasks" --verbose

test: yacc
	go test ${PWD}/pkg/parser

yacc:
	goyacc -o ./pkg/parser/parser.go ./pkg/parser/parser.y
	gofmt -w ./pkg/parser/parser.go

fmt:
	find . -name "*.go" | xargs gofmt -w

deps:
	go get -v all

.PHONY: build parser deps test
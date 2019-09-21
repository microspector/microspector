PWD := $(shell pwd)
export GO111MODULE=on

dist: yacc
	./dist.sh

run: yacc
	go run ./cmd --folder="tasks" --verbose

test: yacc
	go test ${PWD}/pkg/parser

yacc:
	goyacc -xegen ./tasks/main.msf -o ./pkg/parser/parser.go ./pkg/parser/parser.y
	goyacc -xe ./tasks/main.msf -o ./pkg/parser/parser.go ./pkg/parser/parser.y
	gofmt -w ./pkg/parser/parser.go

fmt:
	find . -name "*.go" | xargs gofmt -w

deps:
	go get -v all

.PHONY: build parser deps test
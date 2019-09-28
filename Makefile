PWD := $(shell pwd)
export GO111MODULE=on

dist: yacc
	./dist.sh

run: yacc
	go run ./cmd --file="./tasks/main.msf"

test: yacc
	gofmt -w -s ${PWD}
	go test ${PWD}/pkg/parser

trainyacc:
	goyacc -xegen ./pkg/parser/training.dat -o ./pkg/parser/parser.go ./pkg/parser/parser.y

yacc:
	goyacc -xe ./pkg/parser/training.dat -o ./pkg/parser/parser.go ./pkg/parser/parser.y
	gofmt -w ./pkg/parser/parser.go

fmt:
	find . -name "*.go" | xargs gofmt -w

deps:
	go get -v all

.PHONY: build parser deps test
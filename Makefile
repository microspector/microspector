PWD := $(shell pwd)
export GO111MODULE=on

dist: yacc
	make fmt
	./dist.sh

run:
	go run ./cmd --file="./tasks/main.msf"

test:
	go test ${PWD}/pkg/parser

trainyacc:
	goyacc -xegen ./pkg/parser/training.dat -o ./pkg/parser/parser.go ./pkg/parser/parser.y

yacc:
	goyacc -xe ./pkg/parser/training.dat -o ./pkg/parser/parser.go ./pkg/parser/parser.y

fmt:
	find . -name "*.go" | xargs gofmt -w -s

deps:
	go get -v all

.PHONY: build parser deps test
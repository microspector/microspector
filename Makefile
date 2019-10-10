PWD := $(shell pwd)
export GO111MODULE=on

dist: yacc
	make fmt
	./dist.sh

run:
	go run ./cmd --file="./tasks/main.msf"

runv:
	go run ./cmd --file="./tasks/main.msf" --verbose

test:
	go test ${PWD}/pkg/parser

trainyacc:
	goyacc -xegen ./pkg/parser/training.dat -o ./pkg/parser/parser.go ./pkg/parser/parser.y
	goyacc -xe ./pkg/parser/training.dat -o ./pkg/parser/parser.go ./pkg/parser/parser.y

yacc:
	goyacc  -o ./pkg/parser/parser.go ./pkg/parser/parser.y

fmt:
	find . -name "*.go" | xargs gofmt -w -s

deps:
	go get -v all

.PHONY: build parser deps test
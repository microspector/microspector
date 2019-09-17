package main

import (
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
)

func main() {

	//var file = flag.String("file", "../tasks/main.msf", "Task file path")
	//flag.Parse()

	parser.Parse(`
SET {{ Server }} "cs2744.mojohost.com"
SET {{ Url }} "https://{{ .Server }}"
SET {{ Domain }}  "get.tenta.io"
HTTP GET {{ Url }} HEADER "Host:get.tenta.io" INTO {{ BillingResult }}
DEBUG {{ BillingResult.Error }}
`)

}

package main

import (
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
)

func main() {

	//var file = flag.String("file", "../tasks/main.msf", "Task file path")
	//flag.Parse()

	parser.Parse(`
SET {{ BillingUrl}} "test"
HTTP GET  {{ BillingUrl }} QUERY "username=tufan&password=1"  INTO {{ GetResult }} WHEN {{ BillingUrl}} EQUALS "test2" AND TRUE
HTTP GET  {{ BillingUrl }} HEADER "Host:billing.tenta.io" QUERY "username=tufan&password=1" INTO {{ GetResult }} WHEN TRUE AND TRUE
MUST {{ BillingUrl }} CONTAINS "test"
MUST {{ BillingUrl }} EQUALS "test"
SHOULD {{ BillingUrl }} EQUALS "test" OR {{ BillingUrl }} EQUALS "aa"
`)

}
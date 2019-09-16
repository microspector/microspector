package main

import (
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
)

func main() {

	//var file = flag.String("file", "../tasks/main.msf", "Task file path")
	//flag.Parse()

	parser.Parse(`
SET {{ BillingDomain }} "billing.tenta.io"
SET {{ BillingUrl }} "https://{{ .BillingDomain }}"
DEBUG {{ BillingDomain }} {{ BillingUrl }} WHEN FALSE
HTTP GET {{ BillingUrl }} INTO {{ BillingResult }} 
DEBUG {{ BillingResult.Took }} 
DEBUG {{ BillingResult.Headers.ContentType }} 
DEBUG {{ BillingResult.Headers.XInfra }} 
DEBUG {{ BillingResult.Headers.XPoweredBy }} 
DEBUG "this line will work"
END {{ BillingUrl }} CONTAINS "tentax"
DEBUG "this line will work too"
DEBUG "this line wont work" WHEN FALSE
END WHEN {{ BillingUrl }} CONTAINS "tenta"
DEBUG "this line wont work"
`)

}

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
DEBUG {{ BillingDomain }}
DEBUG {{ BillingUrl }}
HTTP GET {{ BillingUrl }} INTO {{ BillingResult }} 
DEBUG {{ BillingResult.Took }}
DEBUG "this line will work"
END {{ BillingUrl }} CONTAINS "tentax"
DEBUG "this line will work too"
END WHEN {{ BillingUrl }} CONTAINS "tenta"
DEBUG "this line wont work"
`)

}

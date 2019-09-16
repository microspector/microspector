package main

import (
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
)

func main() {

	//var file = flag.String("file", "../tasks/main.msf", "Task file path")
	//flag.Parse()

	parser.Parse(`
SET {{ BillingUrl }} "https://billing.tenta.io" 
HTTP GET  {{ BillingUrl }} 
HTTP GET  {{ BillingUrl }} INTO {{ GetResult }} 
HTTP GET  {{ BillingUrl }} WHEN {{ BillingUrl }} CONTAINS "billing" AND TRUE
HTTP GET  {{ BillingUrl }} INTO {{ GetResult }} WHEN {{ BillingUrl}} EQUALS "test2" AND TRUE
HTTP GET  {{ BillingUrl }} HEADER "Host:billing.tenta.io" QUERY "username=tufan&password=1" INTO {{ GetResult }} WHEN TRUE AND TRUE
MUST {{ BillingUrl }} CONTAINS "test"
MUST {{ BillingUrl }} EQUALS "test"
SHOULD {{ BillingUrl }} EQUALS "test" OR {{ BillingUrl }} EQUALS "aa"
MUST 1 LT 2
DEBUG {{ BillingUrl }}
DEBUG "this line will work"
END {{ BillingUrl }} CONTAINS "tentax"
DEBUG "this line will work too"
END WHEN {{ BillingUrl }} CONTAINS "tenta"
DEBUG "this line wont work"
`)

}
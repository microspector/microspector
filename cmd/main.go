package main

import (
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
)

func main() {

	//var file = flag.String("file", "../tasks/main.msf", "Task file path")
	//flag.Parse()

	parser.Parse(`
SET {{ UserAgent }} "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
SET {{ Server }} "tufans-macbook-pro.local"
SET {{ Url }} "https://{{ .Server }}"
SET {{ Domain }}  "tufans-macbook-pro.local"
HTTP GET {{ Url }} HEADER "Host:{{ .Domain }}\nUser-Agent:{{ .UserAgent }}" INTO {{ BillingResult }}
MUST {{ BillingResult.Json.data._SERVER.HTTP_USER_AGENT }} EQUALS "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
DEBUG {{ BillingResult.Took }} {{ BillingResult.Took }} GT "1000"
`)

}

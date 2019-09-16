package parser

import (
	"net/http"
	"net/url"
)

type HttpCommand struct {
	Method        string
	CommandParams []HttpCommandParam //HEADER, QUERY etc.
	Url           string
}

func (hc *HttpCommand) Run() interface{} {

	u, _ := url.Parse(hc.Url)

	req := &http.Request{
		Method: hc.Method,
		URL:    u,
	}

	client := &http.Client{}
	r, _ := client.Do(req)

	return r
}

type HttpCommandParam struct {
	ParamName  string
	ParamValue string
}

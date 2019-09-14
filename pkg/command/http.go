package command

import (
	"net/http"
	"net/url"
)

//HTTP http_method string_or_var http_command_params INTO variable
type HttpCommand struct {
	Method        string
	CommandParams []HttpCommandParam //HEADER, QUERY etc.
	Url           string
}

func (hc *HttpCommand) run() *http.Response {

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

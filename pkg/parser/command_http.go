package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HttpCommand struct {
	Method        string
	CommandParams []HttpCommandParam //HEADER, QUERY etc.
	Url           string
}

type HttpResult struct {
	Took          int64
	Content       string
	Json          interface{}
	Headers       map[string]string
	StatusCode    int
	ContentLength int
}

func NewFromResponse(response *http.Response, took int64) HttpResult {
	var content []byte
	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := HttpResult{
		ContentLength: len(content),
		Content:       string(content),
		Took:          took,
		Headers:       make(map[string]string),
		StatusCode:    response.StatusCode,
	}

	for k, v := range response.Header {
		result.Headers[toVariableName(k)] = v[0]
	}

	_ = json.Unmarshal(content, &result.Json)

	return result
}

func (hc *HttpCommand) Run() interface{} {

	_, urlError := url.Parse(hc.Url)

	if urlError != nil {
		panic(urlError)
	}

	req, reqError := http.NewRequest(hc.Method, hc.Url, nil)

	if reqError != nil {
		panic(reqError)
	}

	for _, commandParam := range hc.CommandParams {
		switch commandParam.ParamName {
		case "HEADER":
			headers := strings.Split(commandParam.ParamValue, "\n")
			for _, header := range headers {
				headerParts := strings.Split(header, ":")
				if len(headerParts) != 2 {
					panic(fmt.Errorf("error in header format %s", commandParam.ParamValue))
				} else {
					if strings.ToLower(strings.TrimSpace(headerParts[0])) == "host" {
						req.Host = strings.TrimSpace(headerParts[1])
					} else {
						req.Header.Set(strings.TrimSpace(headerParts[0]), strings.TrimSpace(headerParts[1]))
					}
				}

			}

		case "QUERY":

			//how should we format params? query params / post params?

		default:
			fmt.Println("Unknown http command param ", commandParam.ParamName)
		}
	}

	client := &http.Client{}
	start := makeTimestamp()
	r, reqErr := client.Do(req)

	if reqErr != nil {
		panic(reqErr)
	}

	defer r.Body.Close()
	elapsed := makeTimestamp() - start

	return NewFromResponse(r, elapsed)
}

type HttpCommandParam struct {
	ParamName  string
	ParamValue string
}

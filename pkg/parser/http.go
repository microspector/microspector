package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Http struct {
	Method  string
	Url     string
	Params  string
	Into    string
	Headers map[string]string
	Token   Token
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
	}

	//if response.Header.Get("content-type") == "application/json" {
	_ = json.Unmarshal(content, &result.Json)

	//}

	return result
}

var validHttpMethods = map[string]int{
	"GET":     1,
	"POST":    1,
	"HEAD":    1,
	"OPTIONS": 1,
	"PUT":     1,
}

func (h *Http) Run(state *State) error {

	urlToken := h.Token.Tree[2]
	h.Url = urlToken.Text
	if urlToken.isTemplated() {
		_url, err := executeTemplate(urlToken.Text, state.Vars)

		if err != nil {
			log.Println(err)
			return err
		}

		h.Url = _url
	}

	var postForm url.Values

	var resultVariable = "lastCall"

	for i := 0; i < len(h.Token.Tree); i++ {
		token := h.Token.Tree[i]

		switch token.Type {
		case KEYWORD:

			// set parameters here.
			// URL, HEADER, PARAMS, INTO etc.etc.
			if token.Text == "HTTP" {
				//ok
			} else if _, isAMethod := validHttpMethods[token.Text]; isAMethod {
				h.Method = token.Text
			} else if token.Text == "URL" {

				urlToken := h.Token.Tree[i+1]
				h.Url = urlToken.Text

				if urlToken.isTemplated() {
					_url, err := executeTemplate(urlToken.Text, state.Vars)

					if err != nil {
						log.Println(err)
						return err
					}

					h.Url = _url
				}

			} else if token.Text == "PARAMETERS" || token.Text == "PARAMS" {

				paramsToken := h.Token.Tree[i+1]
				h.Params = paramsToken.Text

				if paramsToken.isTemplated() {
					_params, err := executeTemplate(paramsToken.Text, state.Vars)

					if err != nil {
						log.Println(err)
						return err
					}

					h.Params = _params
				}

			} else if token.Text == "HEADER" {
				headersToken := h.Token.Tree[i+1]
				_headers := headersToken.Text

				if headersToken.isTemplated() {
					_evaluatedHeaders, err := executeTemplate(headersToken.Text, state.Vars)

					if err != nil {
						log.Println(err)
						return err
					}
					_headers = _evaluatedHeaders
				}

				headerSegments := strings.Split(_headers, ":")
				if h.Headers == nil {
					h.Headers = map[string]string{}
				}
				log.Printf("Setting Header %s to %s\n", headerSegments[0], headerSegments[1])
				h.Headers[headerSegments[0]] = headerSegments[1]

			} else if token.Text == "INTO" {
				//the variable name.
				h.Into = h.Token.Tree[i+1].Text
			}

			break

		case STRING:
			// skip it for now.
			break

		case VARIABLE:
			// skip it for now
			break

		default:
			log.Printf("Unexpected token : %s", token.TypeName())
			break
		}

	}

	log.Printf("HTTP %s %s\n", h.Method, h.Url)

	r, err := http.NewRequest(h.Method, h.Url, nil)

	if h.Method == "POST" {
		postForm, _ = url.ParseQuery(h.Params)
	} else {
		h.Url = h.Url + "?" + h.Params
	}

	for headerKey, headerValue := range h.Headers {
		if strings.ToUpper(headerKey) == "HOST" {
			r.Host = headerValue
		} else {
			r.Header.Add(headerKey, headerValue)
		}
	}

	//apply headers.
	r.PostForm = postForm

	if err != nil {
		return err
	}

	start := makeTimestamp()
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println(err)
		return err
	}
	elapsed := makeTimestamp() - start

	log.Printf("Got : %d  took : %dms", response.StatusCode, elapsed)

	if err != nil {
		log.Println(err)
		return err
	}

	state.Vars[resultVariable] = NewFromResponse(response, elapsed)

	return nil
}

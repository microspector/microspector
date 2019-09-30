package parser

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpCommand struct {
	Method          string
	CommandParams   []HttpCommandParam //HEADER, BODY etc.
	Url             string
	FollowRedirects bool
}

type HttpResult struct {
	Took          int64
	Content       string
	Json          interface{}
	Headers       map[string]string
	StatusCode    int
	ContentLength int
	Certificate   Certificate
	Error         string
}

type Certificate struct {
	NotAfter int64
}

func NewFromResponse(response *http.Response) HttpResult {
	var content []byte

	if response == nil {
		return HttpResult{}
	}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := HttpResult{
		ContentLength: len(content),
		Content:       string(content),
		Headers:       make(map[string]string),
		StatusCode:    response.StatusCode,
	}

	for k, v := range response.Header {
		result.Headers[ToVariableName(k)] = v[0]
	}

	if response.TLS != nil && len(response.TLS.PeerCertificates) > 0 {
		result.Certificate = Certificate{
			NotAfter: response.TLS.PeerCertificates[0].NotAfter.Unix(),
		}
	}

	_ = json.Unmarshal(content, &result.Json)

	return result
}

func (hc *HttpCommand) Run(l *lex) interface{} {
	_, urlError := url.Parse(hc.Url)

	if urlError != nil {
		panic(urlError)
	}

	req, reqError := http.NewRequest(hc.Method, hc.Url, nil)
	req.Header.Set("User-Agent", fmt.Sprintf("Microspector v%s(%s) (https://microspector.com/ua)", Version, Build))

	if reqError != nil {
		panic(reqError)
	}

	for _, commandParam := range hc.CommandParams {
		switch commandParam.ParamName {
		case "HEADER":
			headers := strings.Split(commandParam.ParamValue.(string), "\n")
			for _, header := range headers {
				headerParts := strings.Split(header, ":")
				if len(headerParts) != 2 {
					panic(fmt.Errorf("error in header format %s", commandParam.ParamValue))
				} else {
					if strings.ToLower(strings.TrimSpace(headerParts[0])) == "host" {
						req.Host = strings.TrimSpace(headerParts[1])
					}
					req.Header.Set(strings.TrimSpace(headerParts[0]), strings.TrimSpace(headerParts[1]))
				}
			}

		case "BODY":
			if req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodPatch {
				req.Body = ioutil.NopCloser(strings.NewReader(commandParam.ParamValue.(string)))
			}
		case "FOLLOW":
			hc.FollowRedirects = IsTrue(commandParam.ParamValue)
		default:
			fmt.Println("Unknown http command param ", commandParam.ParamName)
		}
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	dialDepth := 1
	conf := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         req.Host,
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: conf,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				// redirect all connections to host specified in url
				if dialDepth == 1 { // use the host in url just for firs dial, others
					addr = strings.Split(req.URL.Host, ":")[0] + addr[strings.LastIndex(addr, ":"):]
					dialDepth++
				} else {
					conf.ServerName = ""
				}
				return dialer.DialContext(ctx, network, addr)
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if !hc.FollowRedirects {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
	start := makeTimestamp()
	r, reqErr := client.Do(req)
	if reqErr == nil {
		defer r.Body.Close()
	}
	elapsed := makeTimestamp() - start

	if r.TLS != nil {
		fmt.Println("r.TLS.PeerCertificates[0].NotAfter", r.TLS.PeerCertificates[0].NotAfter)
	}

	resp := NewFromResponse(r)
	resp.Took = elapsed
	if reqErr != nil {
		resp.Error = reqErr.Error()
	}

	return resp
}

type HttpCommandParam struct {
	ParamName  string
	ParamValue interface{}
}

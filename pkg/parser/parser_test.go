package parser

import (
	"fmt"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest() *httptest.Server {
	serverMux := http.NewServeMux()
	server := httptest.NewServer(serverMux)

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Microspector")
		w.Header().Set("Microspector", "Service Up")

	})

	return server
}
func TestParser_Set(t *testing.T) {

	lex := Parse(`
SET {{ Domain }} "microspector.com"
SET {{ ContainsTrue }}  "microspector.com" CONTAINS "microspector"
SET {{ ContainsFalse }} "microspector.com" CONTAINS "microspectorFAIL"
SET {{ StartsWithTrue }} "microspector.com" STARTSWITH "microspector"
SET {{ StartsWithFalse }} "microspector.com" STARTSWITH "microspectorFAIL"
SET {{ DoubleDomain }} "microspector.com {{ .Domain }}"
SET {{ Hundred }} 100
SET {{ StringDigitCompare }} "100" LT 101
SET {{ StringDigitCompare2 }} "100" GT 99
SET {{ StringDigitCompare3 }} 100 GT "99"
SET {{ StringDigitCompare4 }} {{ Hundred }} GT "99"
SET {{ StringDigitCompare5 }} {{ Hundred }} GT "999"
SET {{ WhenFalse }} FALSE WHEN "100" < "101"
	`)

	Run(lex)

	assert.Equal(t, GlobalVars["Domain"], "microspector.com")
	assert.Equal(t, GlobalVars["ContainsTrue"], true)
	assert.Equal(t, GlobalVars["ContainsFalse"], false)
	assert.Equal(t, GlobalVars["StartsWithTrue"], true)
	assert.Equal(t, GlobalVars["StartsWithFalse"], false)
	assert.Equal(t, GlobalVars["DoubleDomain"], "microspector.com microspector.com")
	assert.Equal(t, GlobalVars["StringDigitCompare"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare2"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare3"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare4"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare5"], false)
	assert.Equal(t, GlobalVars["WhenFalse"], false)
}

func TestParser_Http(t *testing.T) {
	server := setupTest()
	defer server.Close()

	GlobalVars["ServerMux"] = server.URL

	Run(Parse(`
HTTP GET {{ ServerMux }} INTO {{ ServerResult }}
SET {{ ContentLength }} {{ ServerResult.ContentLength }}
	`))

	assert.Equal(t, GlobalVars["ServerMux"], server.URL)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).ContentLength, 12)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).StatusCode, 200)

}

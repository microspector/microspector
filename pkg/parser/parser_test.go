package parser

import (
	"fmt"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest() *http.ServeMux {
	serverMux := http.NewServeMux()
	server := httptest.NewServer(serverMux)
	defer server.Close()

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Host", r.Host)
		fmt.Fprintln(w, r.Header)

	})

	return serverMux
}
func TestParser_Set(t *testing.T) {

	Run(Parse(`
SET {{ Domain }} "microspector.com"
SET {{ ContainsTrue }} "microspector.com" CONTAINS "microspector"
SET {{ ContainsFalse }} "microspector.com" CONTAINS "microspectorFAIL"
SET {{ StartsWithTrue }} "microspector.com" CONTAINS "microspector"
SET {{ StartsWithFalse }} "microspector.com" CONTAINS "microspectorFAIL"
SET {{ DoubleDomain }} "microspector.com {{ .Domain }}"
SET {{ Hundred }} 100
SET {{ StringDigitCompare }} "100" LT 101
SET {{ StringDigitCompare2 }} "100" GT 99
SET {{ StringDigitCompare3 }} 100 GT "99"
SET {{ StringDigitCompare4 }} {{ Hundred }} GT "99"
	`))
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
}

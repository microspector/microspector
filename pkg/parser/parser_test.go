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
		w.Header().Set("Microspector", "Service Up")
		w.Header().Set("User-Agent", r.Header.Get("User-Agent"))
		w.Header().Set("Host", r.Header.Get("Host"))
		fmt.Fprint(w, `{"data":"microspector.com"}`)

	})

	return server
}
func TestParser_Set(t *testing.T) {

	lex := Parse(`
SET {{ Domain }} 'microspector.com'
SET {{ ContainsTrue }}  "microspector.com" contains "microspector"
SET {{ ContainsFalse }} "microspector.com" CONTAINS "microspectorFAIL"
SET {{ StartsWithTrue }} "microspector.com" startswith "microspector"
SET {{ StartsWithFalse }} "microspector.com" STARTSWITH "microspectorFAIL"
SET {{ DoubleDomain }} "microspector.com {{ .Domain }}"
SET {{ Hundred }} 100
SET {{ StringDigitCompare1 }} '100' LT 101 AND 100 == 100
SET {{ StringDigitCompare2 }} ("100" < 101 AND "100" equals 20 * 5) OR (1 != 1)
SET {{ StringDigitCompare3 }} "100" <= 101
SET {{ StringDigitCompare4 }} "100" GT 99
SET {{ StringDigitCompare5 }} "100" > 99
SET {{ StringDigitCompare6 }} "100" >= 99
SET {{ StringDigitCompare7 }} "100" <= 99
SET {{ StringDigitCompare8 }} "100" == 100
SET {{ StringDigitCompare9 }} "100" != 100
SET {{ StringDigitCompare10 }} 100 GT "99"
SET {{ StringDigitCompare11 }} {{ Hundred }} gt "99"
SET {{ StringDigitCompare12 }} {{ Hundred }} GT "999"
SET {{ WhenFalse }} false when "100" < "101"
SET {{ SSLRand }} "{{ openssl_rand 32 \"hex\" }}"
SET {{ SSLRandSize }} "{{ str_len .SSLRand }}"
SET {{ HashMd5 }} "{{ hash_md5 \"1\" }}"
SET {{ Hash256 }} "{{ hash_sha256 .HashMd5 }}"
`)

	Run(lex)

	assert.Equal(t, GlobalVars["Domain"], "microspector.com")
	assert.Equal(t, GlobalVars["ContainsTrue"], true)
	assert.Equal(t, GlobalVars["ContainsFalse"], false)
	assert.Equal(t, GlobalVars["StartsWithTrue"], true)
	assert.Equal(t, GlobalVars["StartsWithFalse"], false)
	assert.Equal(t, GlobalVars["DoubleDomain"], "microspector.com microspector.com")
	assert.Equal(t, GlobalVars["StringDigitCompare1"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare2"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare3"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare4"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare5"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare6"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare7"], false)
	assert.Equal(t, GlobalVars["StringDigitCompare8"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare9"], false)
	assert.Equal(t, GlobalVars["StringDigitCompare10"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare11"], true)
	assert.Equal(t, GlobalVars["StringDigitCompare12"], false)
	assert.Equal(t, GlobalVars["WhenFalse"], false)
	assert.Equal(t, GlobalVars["SSLRandSize"], "64")
	assert.Equal(t, len(GlobalVars["HashMd5"].(string)), 32)
	assert.Equal(t, GlobalVars["HashMd5"].(string), "c4ca4238a0b923820dcc509a6f75849b")
	assert.Equal(t, len(GlobalVars["Hash256"].(string)), 64)
	assert.Equal(t, GlobalVars["Hash256"].(string), "08428467285068b426356b9b0d0ae1e80378d9137d5e559e5f8377dbd6dde29f")
}

func TestParser_Http(t *testing.T) {
	server := setupTest()
	defer server.Close()

	GlobalVars["ServerMux"] = server.URL

	Run(Parse(`
HTTP get {{ ServerMux }} HEADER "User-Agent:(bot)microspector.com" INTO {{ ServerResult }}
SET {{ ContentLength }} {{ ServerResult.ContentLength }}
set {{ RawContent }} {{ ServerResult.Content }}
SET {{ ContentData }} {{ ServerResult.Json.data }}
	`))

	assert.Equal(t, GlobalVars["ServerMux"], server.URL)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).ContentLength, 27)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).StatusCode, 200)
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).Headers["UserAgent"], "(bot)microspector.com")
	assert.Equal(t, GlobalVars["ServerResult"].(HttpResult).Headers["Microspector"], "Service Up")
	assert.Equal(t, GlobalVars["RawContent"], `{"data":"microspector.com"}`)
	assert.Equal(t, GlobalVars["ContentData"], "microspector.com")

}

func TestParser_End(t *testing.T) {
	lex := Parse(`
SET {{ Var50 }} 49
END WHEN {{ Var50 }} > 100 #this line won't end the execution
SET {{ Var50 }} 50
end
SET {{ Var50 }} 100
SET {{ NilVar }} "this should not be assigned"
`)

	Run(lex)

	assert.Equal(t, GlobalVars["Var50"], int64(50))
	assert.Equal(t, GlobalVars["NilVar"], nil)
}

func TestParser_Assert(t *testing.T) {
	lex := Parse(`
SET {{ Var50 }} 50
ASSERT {{ Var50 }} > 100 
ASSERT {{ Var50 }} < 100 
ASSERT {{ Var50 }} < 100 
SHOULD {{ Var50 }} < 100 
SHOULD {{ Var50 }} > 100 
MUST {{ Var50 }} > 100 
MUST {{ Var50 }} < 100 
`)

	Run(lex)

	assert.Equal(t, State.Assertion.Failed, 1)
	assert.Equal(t, State.Assertion.Succeeded, 2)

	assert.Equal(t, State.Should.Failed, 1)
	assert.Equal(t, State.Should.Succeeded, 1)

	assert.Equal(t, State.Must.Failed, 1)
	assert.Equal(t, State.Must.Succeeded, 1)

}

func TestParser_QuotedString(t *testing.T) {
	lex := Parse(`
SET {{ SingleQuoted50 }} '50'
SET {{ DoubleQuoted50 }} "50"
SET {{ SingleContainsDouble }}  '"this is a string "" with double quotes"'
SET {{ DoubleContainsSingle }}  "'this is a string '' with single quotes'"
SET {{ SingleContainsSingle }}  '\'this is a string \'\' with single includes quotes\''
SET {{ DoubleContainsDouble }}  "\"this is a string \"\" with double includes quotes\""
SET {{ DoubleContainsDouble }}  "\"this is a string \"\" with double includes quotes\""
` + " SET {{ BackTicks }} `\"\"this is' a back tick yeah\"\"` ")

	Run(lex)

	assert.Equal(t, GlobalVars["SingleQuoted50"], "50")
	assert.Equal(t, GlobalVars["DoubleQuoted50"], "50")
	assert.Equal(t, GlobalVars["SingleContainsDouble"], `"this is a string "" with double quotes"`)
	assert.Equal(t, GlobalVars["DoubleContainsSingle"], `'this is a string '' with single quotes'`)
	assert.Equal(t, GlobalVars["SingleContainsSingle"], `'this is a string '' with single includes quotes'`)
	assert.Equal(t, GlobalVars["DoubleContainsDouble"], `"this is a string "" with double includes quotes"`)
	assert.Equal(t, GlobalVars["BackTicks"], `""this is' a back tick yeah""`)

}

func TestParser_Arithmetic(t *testing.T) {

	lex := Parse(`
SET {{ Var1 }} 1
SET {{ Var2 }} '2' # it should support both
SET {{ Var3 }} 3
SET {{ Var4 }} 4
SET {{ Var5 }} '5'
SET {{ Var6 }} 6
SET {{ Var7 }} 7
SET {{ Var8 }} '8'
SET {{ Var9 }} 9
SET {{ Var10 }} 10

SET {{ Result10 }} {{ Var5 }} * 2
SET {{ Result15 }} {{ Var3 }} * 5
SET {{ Result5 }} {{ Var10 }} / 2
SET {{ Result10Strings }} {{ Var8 }} + {{ Var2 }}
SET {{ Result6Strings }} {{ Var8 }} - {{ Var2 }}
SET {{ Result5Strings }} {{ Var8 }} - {{ Var2 }} - 1
SET {{ Result501Strings }} {{ Var10 }} * {{ Var5 }} * 5 + 1 + {{ Var10 }} * {{ Var5 }} * 5
SET {{ Result550Strings }} {{ Var10 }} * {{ Var5 }} * (5 + 1) + {{ Var10 }} * {{ Var5 }} * 5
SET {{ Result262Strings }} {{ Var10 }} / {{ Var5 }} * (5 + 1) + {{ Var10 }} * {{ Var5 }} * 5
SET {{ ResultFloat15 }} {{ Var10 }} * 1.5
SET {{ ResultFloat10 }} {{ ResultFloat15 }} / 1.5
`)

	Run(lex)

	assert.Equal(t, GlobalVars["Var1"], int64(1))
	assert.Equal(t, GlobalVars["Var2"], "2")
	assert.Equal(t, GlobalVars["Var3"], int64(3))
	assert.Equal(t, GlobalVars["Var4"], int64(4))
	assert.Equal(t, GlobalVars["Var5"], "5")
	assert.Equal(t, GlobalVars["Var6"], int64(6))
	assert.Equal(t, GlobalVars["Var7"], int64(7))
	assert.Equal(t, GlobalVars["Var8"], "8")
	assert.Equal(t, GlobalVars["Var9"], int64(9))
	assert.Equal(t, GlobalVars["Var10"], int64(10))

	assert.Equal(t, GlobalVars["Result10"], int64(10))
	assert.Equal(t, GlobalVars["Result15"], float64(15))
	assert.Equal(t, GlobalVars["Result5"], float64(5))
	assert.Equal(t, GlobalVars["Result6Strings"], int64(6))
	assert.Equal(t, GlobalVars["Result5Strings"], int64(5))
	assert.Equal(t, GlobalVars["Result501Strings"], float64(501))
	assert.Equal(t, GlobalVars["Result550Strings"], float64(550))
	assert.Equal(t, GlobalVars["Result262Strings"], float64(262))
	assert.Equal(t, GlobalVars["ResultFloat15"], float64(15))
	assert.Equal(t, GlobalVars["ResultFloat10"], float64(10))

}

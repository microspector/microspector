package parser

import (
	"fmt"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

	serverMux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", server.URL+"/redirected")
		w.WriteHeader(302)
	})

	serverMux.HandleFunc("/redirected", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `OK`)
	})

	serverMux.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Second)
		fmt.Fprint(w, `OK`)
	})

	return server
}

func TestParser_Http(t *testing.T) {

	server := setupTest()
	defer server.Close()

	l := Parse(`
HTTP get {{ ServerMux }} HEADER "User-Agent:(bot)microspector.com" INTO {{ ServerResult }}
SET {{ ContentLength }} {{ ServerResult.ContentLength }}
set {{ RawContent }} {{ ServerResult.Content }}
SET {{ ContentData }}  ServerResult.Json.data 
MUST $ContentData equals "microspector.com"
MUST ServerResult.Json.data  equals "microspector.com"
	`)

	l.GlobalVars["ServerMux"] = server.URL

	Run(l)

	assert.Equal(t, l.GlobalVars["ServerMux"], server.URL)
	assert.Equal(t, l.GlobalVars["ServerResult"].(HttpResult).ContentLength, 27)
	assert.Equal(t, l.GlobalVars["ServerResult"].(HttpResult).StatusCode, 200)
	assert.Equal(t, l.GlobalVars["ServerResult"].(HttpResult).Headers["UserAgent"], "(bot)microspector.com")
	assert.Equal(t, l.GlobalVars["ServerResult"].(HttpResult).Headers["Microspector"], "Service Up")
	assert.Equal(t, l.GlobalVars["RawContent"], `{"data":"microspector.com"}`)
	assert.Equal(t, l.GlobalVars["ContentData"], "microspector.com")
	assert.Equal(t, l.State.Must.Succeeded, 2)

}

func TestParser_HttpRedirect(t *testing.T) {

	server := setupTest()
	defer server.Close()

	l := Parse(`
HTTP get "{{ .ServerMux }}/redirect" HEADER "User-Agent:(bot)microspector.com" NOFOLLOW INTO {{ ServerResult }}
MUST ServerResult.StatusCode == 302
HTTP get "{{ .ServerMux }}/redirect" HEADER "User-Agent:(bot)microspector.com" FOLLOW INTO {{ ServerResultNoFollow }}
MUST ServerResultNoFollow.StatusCode == 200
	`)

	l.GlobalVars["ServerMux"] = server.URL

	Run(l)

	assert.Equal(t, l.GlobalVars["ServerMux"], server.URL)
	assert.Equal(t, l.GlobalVars["ServerResult"].(HttpResult).Headers["Location"], server.URL+"/redirected")
	assert.Equal(t, l.GlobalVars["ServerResultNoFollow"].(HttpResult).Content, "OK")
	assert.Equal(t, l.State.Must.Succeeded, 2)

}

func TestParser_Set(t *testing.T) {

	lex := Parse(`
SET {{ Domain }} 'microspector.com'
SET $ContainsTrue  "microspector.com" contains "microspector"
SET $ContainsFalse  "microspector.com" CONTAINS "microspectorFAIL"
SET $StartsWithTrue  "microspector.com" startswith "microspector"
SET $StartsWithFalse  "microspector.com" STARTSWITH "microspectorFAIL"
SET $DoubleDomain "microspector.com {{ .Domain }}"
SET {{ Hundred }} 100
SET {{ StringDigitCompare1 }} '100' LT 101 AND 100 == 100
SET {{ StringDigitCompare2 }} ("100" < 101 AND "100" equals 20 * 5) OR (1 != 1)
SET {{ StringDigitCompare3 }} "100" <= 101
SET {{ StringDigitCompare4 }} "100" GT 99
SET {{ StringDigitCompare5 }} "100" > 99
SET {{ StringDigitCompare6 }} "100" >= 99
SET {{ StringDigitCompare7 }} "100" <= 99
SET {{ StringDigitCompare8 }} "100" == ( 200 / 4 ) * 2
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

	assert.Equal(t, lex.GlobalVars["Domain"], "microspector.com")
	assert.Equal(t, lex.GlobalVars["ContainsTrue"], true)
	assert.Equal(t, lex.GlobalVars["ContainsFalse"], false)
	assert.Equal(t, lex.GlobalVars["StartsWithTrue"], true)
	assert.Equal(t, lex.GlobalVars["StartsWithFalse"], false)
	assert.Equal(t, lex.GlobalVars["DoubleDomain"], "microspector.com microspector.com")
	assert.Equal(t, lex.GlobalVars["StringDigitCompare1"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare2"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare3"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare4"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare5"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare6"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare7"], false)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare8"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare9"], false)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare10"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare11"], true)
	assert.Equal(t, lex.GlobalVars["StringDigitCompare12"], false)
	assert.Equal(t, lex.GlobalVars["WhenFalse"], false)
	assert.Equal(t, lex.GlobalVars["SSLRandSize"], "64")
	assert.Equal(t, len(lex.GlobalVars["HashMd5"].(string)), 32)
	assert.Equal(t, lex.GlobalVars["HashMd5"].(string), "c4ca4238a0b923820dcc509a6f75849b")
	assert.Equal(t, len(lex.GlobalVars["Hash256"].(string)), 64)
	assert.Equal(t, lex.GlobalVars["Hash256"].(string), "08428467285068b426356b9b0d0ae1e80378d9137d5e559e5f8377dbd6dde29f")
}

func TestParser_End(t *testing.T) {

	lex := Parse(`
SET {{ Var50 }} 49
END WHEN {{ Var50 }} > 100 #this line won't end the execution
SET {{ Var50 }} 50
SET varNegative  -1 
SET negativeResult varNegative * Var50
end
SET {{ Var50 }} 100
SET {{ NilVar }} "this should not be assigned"
`)

	Run(lex)

	assert.Equal(t, lex.GlobalVars["Var50"], int64(50))
	assert.Equal(t, lex.GlobalVars["NilVar"], nil)
	assert.Equal(t, lex.GlobalVars["varNegative"], int64(-1))
	assert.Equal(t, lex.GlobalVars["negativeResult"], int64(-50))
}

func TestParser_Must(t *testing.T) {

	l := Parse(`
SET {{ Var50 }} 49
SET {{ VarTrue }} true
SET {{ VarFalse }} true AND false
MUST "{{ .VarFalse }}" EQUALS "false"
MUST {{ VarFalse }} EQUALS false
MUST {{ Var50 }} EQUALS 50
MUST {{ Var50 }} EQUALS 49
SET MyVar "microspector"
SHOULD $MyVar contain "micro"
SHOULD $MyVar contain "amicro"
ASSERT $MyVar contain "amicro"
ASSERT $MyVar contain "micro"
`)

	Run(l)
	//TODO: do some more assertion, like, must fail and success counts
	assert.Equal(t, l.GlobalVars["VarFalse"], false)
	assert.Equal(t, l.State.Must.Failed, 1)
	assert.Equal(t, l.State.Must.Succeeded, 3)
	assert.Equal(t, l.State.Should.Succeeded, 1)
	assert.Equal(t, l.State.Should.Failed, 1)
	assert.Equal(t, l.State.Assert.Failed, 1)
	assert.Equal(t, l.State.Assert.Succeeded, 1)
}

func TestParser_Assert(t *testing.T) {

	l := Parse(`
SET {{ Var50 }} 50
ASSERT {{ Var50 }} > 100 
ASSERT {{ Var50 }} < 100 
ASSERT {{ Var50 }} < 100 
SHOULD {{ Var50 }} < 100 
SHOULD {{ Var50 }} > 100 
MUST {{ Var50 }} > 100 
MUST {{ Var50 }} < 100 
`)

	Run(l)

	assert.Equal(t, l.State.Assert.Failed, 1)
	assert.Equal(t, l.State.Assert.Succeeded, 2)

	assert.Equal(t, l.State.Should.Failed, 1)
	assert.Equal(t, l.State.Should.Succeeded, 1)

	assert.Equal(t, l.State.Must.Failed, 1)
	assert.Equal(t, l.State.Must.Succeeded, 1)

}

func TestParser_QuotedString(t *testing.T) {

	l := Parse(`
SET {{ SingleQuoted50 }} '50'
SET {{ DoubleQuoted50 }} "50"
SET {{ SingleContainsDouble }}  '"this is a string "" with double quotes"'
SET {{ DoubleContainsSingle }}  "'this is a string '' with single quotes'"
SET {{ SingleContainsSingle }}  '\'this is a string \'\' with single includes quotes\''
SET {{ DoubleContainsDouble }}  "\"this is a string \"\" with double includes quotes\""
SET {{ DoubleContainsDouble }}  "\"this is a string \"\" with double includes quotes\""
` + " SET {{ BackTicks }} `\"\"this is' a back tick yeah\"\"` ")

	Run(l)

	assert.Equal(t, l.GlobalVars["SingleQuoted50"], "50")
	assert.Equal(t, l.GlobalVars["DoubleQuoted50"], "50")
	assert.Equal(t, l.GlobalVars["SingleContainsDouble"], `"this is a string "" with double quotes"`)
	assert.Equal(t, l.GlobalVars["DoubleContainsSingle"], `'this is a string '' with single quotes'`)
	assert.Equal(t, l.GlobalVars["SingleContainsSingle"], `'this is a string '' with single includes quotes'`)
	assert.Equal(t, l.GlobalVars["DoubleContainsDouble"], `"this is a string "" with double includes quotes"`)
	assert.Equal(t, l.GlobalVars["BackTicks"], `""this is' a back tick yeah""`)

}

func TestParser_Arithmetic(t *testing.T) {

	l := Parse(`
SET  Var1  1
SET {{ Var2 }} '2' # it should support both
SET {{ Var3 }} 3
SET {{ Var4 }} 4
SET {{ Var5 }} '5'
SET {{ Var6 }} 6
SET {{ Var7 }} 7
SET {{ Var8 }} '8'
SET {{ Var9 }} 9
SET {{ Var10 }} 10

SET {{ Result10 }} Var5 * 2
SET {{ Result15 }} {{ Var3 }} * 5
SET {{ Result5 }} {{ Var10 }} / 2
SET {{ Result10Strings }} {{ Var8 }} + {{ Var2 }}
SET {{ Result6Strings }} {{ Var8 }} - {{ Var2 }}
SET {{ Result5Strings }} {{ Var8 }} - {{ Var2 }} - 1
SET {{ Result501Strings }} {{ Var10 }} * {{ Var5 }} * 5 + 1 + {{ Var10 }} * {{ Var5 }} * 5
SET {{ Result550Strings }} {{ Var10 }} * {{ Var5 }} * (5 + 1) + {{ Var10 }} * {{ Var5 }} * 5
SET {{ Result262Strings }} {{ Var10 }} / {{ Var5 }} * (5 + 1) + {{ Var10 }} * {{ Var5 }} * 5
SET {{ ResultFloat15 }} {{ Var10 }} * 1.5
SET $ResultFloat10 {{ ResultFloat15 }} / 1.5
`)

	Run(l)

	assert.Equal(t, l.GlobalVars["Var1"], int64(1))
	assert.Equal(t, l.GlobalVars["Var2"], "2")
	assert.Equal(t, l.GlobalVars["Var3"], int64(3))
	assert.Equal(t, l.GlobalVars["Var4"], int64(4))
	assert.Equal(t, l.GlobalVars["Var5"], "5")
	assert.Equal(t, l.GlobalVars["Var6"], int64(6))
	assert.Equal(t, l.GlobalVars["Var7"], int64(7))
	assert.Equal(t, l.GlobalVars["Var8"], "8")
	assert.Equal(t, l.GlobalVars["Var9"], int64(9))
	assert.Equal(t, l.GlobalVars["Var10"], int64(10))

	assert.Equal(t, l.GlobalVars["Result10"], int64(10))
	assert.Equal(t, l.GlobalVars["Result15"], int64(15))
	assert.Equal(t, l.GlobalVars["Result5"], int64(5))
	assert.Equal(t, l.GlobalVars["Result6Strings"], int64(6))
	assert.Equal(t, l.GlobalVars["Result5Strings"], int64(5))
	assert.Equal(t, l.GlobalVars["Result501Strings"], int64(501))
	assert.Equal(t, l.GlobalVars["Result550Strings"], int64(550))
	assert.Equal(t, l.GlobalVars["Result262Strings"], int64(262))
	assert.Equal(t, l.GlobalVars["ResultFloat15"], float64(15))
	assert.Equal(t, l.GlobalVars["ResultFloat10"], float64(10))

}

/**
https://twitter.com/s0md3v/status/1171394403065155584

1. cat matches cat
2. ca+t matches caaaaaaaaaaaat but not ct
3. ca*t matches caaaaaaaaaaaat and also ct
4. ca{2,4} matches caat, caaat and caaaat
5. c(at)+ matches catatatatatat
6. c(at|orn) matches cat and corn
7. c[ea]t matches cat and cet
8. c[ea]+t matches caaaat and ceeet
9. c[A-C0-9]t matches cAt, cBt, cCt, c8t etc.
10. c.t matches cat, c&t, c2t (any char between c and t)
11. c.+t matches c3%x4t (any number of any chars)
12. c.*t matches c3%x4t and as well as ct
13. ^ denotes start of a string, $ denotes the end
14. ^a+cat will match aaacat in aaacat but not in bbaaacat
15. cat$ will match cat in aaacat but not in aaacats
16. ^cat$ will match only and only this string i.e. cat

\d is for digits, \w for alphanumeric chars, \s is for white space chars & line breaks
\D is for non-digits, \W for non-alphamueric chars and \s is for non-white space chars
\t for tabs, \r for carriage return and \n for newline

Yes, c\d+t matches c2784t
Yes, c\s+ matches c       t
Yes, c\D+ matches cxxxt ca2t

Using .*w vs .*?w on xxsomethingnew@1234wxx
.*w returns somethingnew@1234w (longest match)
.*w? returns somethingnew (shortest match)
*/
func TestParser_Regex(t *testing.T) {

	l := Parse(`
SET {{ Cat }} "cat"
SET Regex1 {{ Cat }} MATCHES "cat"
SET {{ Regex2 }} "ca+t" MATCHES "caaaaaaaaaaaat"
SET {{ Regex2a }} "ca+t" MATCHES "ct"
SET {{ Regex3 }} "ca*t" MATCHES "caaaaaaaaaaaat"
SET {{ Regex3a }} "ca*t" MATCHES "ct"
SET {{ Regex4a }} "ca{2,4}" MATCHES "caaaat"
SET {{ Regex4b }} "ca{2,4}" MATCHES {{ Cat }}
SET {{ Regex4c }} "ca{2,4}" MATCHES "caaaaat"
SET {{ Regex5 }} "c(at)+" MATCHES "catatatatatat"
SET {{ Regex6 }} "c(at|orn)" MATCHES "cat" AND "c(at|orn)" MATCHES "cat"
SET {{ Regex7 }} "c[ea]t" MATCHES "cat" AND "c[ea]t" MATCHES "cet"
SET {{ Regex8 }} "c[ea]+t" MATCHES  "caaaat" AND "c[ea]+t" MATCHES "ceeet"   # c[ea]+t matche caaaat and ceeet
SET {{ Regex9 }} "c[A-C0-9]t" MATCHES "cAt" AND "c[A-C0-9]t" MATCHES "c8t"   # c[A-C0-9]t matche cAt, cBt, cCt, c8t etc.
SET {{ Regex10 }} "c.t" MATCHES "cat" AND "c.t" MATCHES "c&t" # c.t matche cat, c&t, c2t (any char between c and t)
SET {{ Regex11 }} "c.+t" MATCHES "c3%x4t"  # c.+t matche c3%x4t (any number of any chars)
SET {{ Regex12 }} "c.*t" MATCHES "c3%x4t" # c.*t matche c3%x4t and as well as ct
SET {{ Regex13 }} true # skip this :)
SET {{ Regex14 }} "^a+cat" MATCHES "aaacat" # ^a+cat will match aaacat in aaacat but not in bbaaacat
SET {{ Regex14a }} "^a+cat" MATCHES "bbaaacat" # ^a+cat will match aaacat in aaacat but not in bbaaacat
SET {{ Regex15 }} "cat$" MATCHES  "aaacat" # cat$ will match cat in aaacat but not in aaacats
SET {{ Regex15a }} "cat$" MATCHES  "aaacats" # cat$ will match cat in aaacat but not in aaacats
SET {{ Regex16 }} "^cat$" MATCHES "cat" # ^cat$ will match only and only this string i.e. cat
SET {{ Regex16a }} "^cat$" MATCHES "cata" # ^cat$ will match only and only this string i.e. cat
`)

	Run(l)

	assert.Equal(t, l.GlobalVars["Regex1"], true)
	assert.Equal(t, l.GlobalVars["Regex2"], true)
	assert.Equal(t, l.GlobalVars["Regex2a"], false)
	assert.Equal(t, l.GlobalVars["Regex3"], true)
	assert.Equal(t, l.GlobalVars["Regex3a"], true)
	assert.Equal(t, l.GlobalVars["Regex4a"], true)
	assert.Equal(t, l.GlobalVars["Regex4b"], false)
	assert.Equal(t, l.GlobalVars["Regex4c"], true)
	assert.Equal(t, l.GlobalVars["Regex5"], true)
	assert.Equal(t, l.GlobalVars["Regex6"], true)
	assert.Equal(t, l.GlobalVars["Regex7"], true)
	assert.Equal(t, l.GlobalVars["Regex8"], true)
	assert.Equal(t, l.GlobalVars["Regex9"], true)
	assert.Equal(t, l.GlobalVars["Regex10"], true)
	assert.Equal(t, l.GlobalVars["Regex11"], true)
	assert.Equal(t, l.GlobalVars["Regex12"], true)
	assert.Equal(t, l.GlobalVars["Regex13"], true)
	assert.Equal(t, l.GlobalVars["Regex14"], true)
	assert.Equal(t, l.GlobalVars["Regex14a"], false)
	assert.Equal(t, l.GlobalVars["Regex15"], true)
	assert.Equal(t, l.GlobalVars["Regex15a"], false)
	assert.Equal(t, l.GlobalVars["Regex16"], true)
	assert.Equal(t, l.GlobalVars["Regex16a"], false)

}

func TestParser_TypeCheck(t *testing.T) {

	l := Parse(`
SET {{ String }} "cat"
MUST {{ String }} is string
SET {{ Boolean }} String is string
MUST Boolean is boolean
MUST Boolean is bool
SET {{ Integer }} 10
MUST {{ Integer }} is int ## known issue: https://golang.org/pkg/encoding/json/#Unmarshal
MUST {{ Integer }} is integer ## known issue: https://golang.org/pkg/encoding/json/#Unmarshal
`)

	Run(l)
	assert.Equal(t, l.State.Must.Succeeded, 5)
}

func TestParser_Arrays(t *testing.T) {

	l := Parse(`
SET {{ String }} "onestring"
SET {{ Array }} ["cat",1,"2",1.0,{{ String }},[1,2,3,4]]
`)

	Run(l)

	assert.Equal(t, len(l.GlobalVars["Array"].([]interface{})), 6)
	assert.Equal(t, l.GlobalVars["Array"].([]interface{})[4], "onestring") //5th element of array should be $String which is "onestring"
	assert.DeepEqual(t, l.GlobalVars["Array"].([]interface{})[5], []interface{}{int64(1), int64(2), int64(3), int64(4)})
}

func TestParser_NotOperator(t *testing.T) {

	l := Parse(`
SET {{ Bool }} true
SET $Microspector "Microspector"
MUST {{ Bool }} not equals false
MUST {{ Bool }} equals true
MUST $Microspector CONTAIN "Microspector"
MUST $Microspector NOT contain "mcrospector"
MUST $Microspector NOT contains "mcrospector"
MUST NOT $Microspector contains "mcrospector"
SHOULD NOT $Microspector contains "mcrospector"
Set $age 30
must not $age < 10
must age not < 10
must age > 30
`)

	Run(l)

	assert.Equal(t, l.GlobalVars["Bool"], true)
	assert.Equal(t, l.State.Must.Succeeded, 8)
	assert.Equal(t, l.State.Must.Failed, 1)
	assert.Equal(t, l.State.Should.Succeeded, 1)
}

func TestParser_Funcs(t *testing.T) {

	l := Parse(`
set length20 len("tufan baris yildirim")
set length64 len(openssl_rand(32,"hex"))
must len(openssl_rand(32,"hex")) == 64
must 64 == len(openssl_rand(32,"hex"))
`)

	Run(l)

	assert.Equal(t, l.GlobalVars["length20"], 20)
	assert.Equal(t, l.GlobalVars["length64"], 64)
	assert.Equal(t, l.State.Must.Succeeded, 2)
}

func TestParser_Cmd(t *testing.T) {

	l := Parse(`
set length20 len("tufan baris yildirim")
set echo "echo"
cmd echo "{{ .length20 }}" into x
cmd echo $length20 into x2
cmd 'echo' 'microspector' into output
must output equals 'microspector'
`)

	Run(l)

	assert.Equal(t, l.GlobalVars["length20"], 20)
	assert.Equal(t, l.GlobalVars["x"], "20")
	assert.Equal(t, l.GlobalVars["x2"], "20")
	assert.Equal(t, l.State.Must.Succeeded, 1)
}

func TestParser_In(t *testing.T) {

	l := Parse(`
SET {{ String }} "onestring"
SET {{ Array }} ["cat",1,"2",1.0,{{ String }},[1,2,3,4]]
MUST "onestring" IN {{ Array }}
MUST "onestring2" IN {{ Array }}
MUST NOT "onestring2" IN {{ Array }}
`)

	Run(l)

	assert.Equal(t, len(l.GlobalVars["Array"].([]interface{})), 6)
	assert.Equal(t, l.GlobalVars["Array"].([]interface{})[4], "onestring") //5th element of array should be $String which is "onestring"
	assert.DeepEqual(t, l.GlobalVars["Array"].([]interface{})[5], []interface{}{int64(1), int64(2), int64(3), int64(4)})
	assert.Equal(t, l.State.Must.Succeeded, 2)
	assert.Equal(t, l.State.Must.Failed, 1)
}

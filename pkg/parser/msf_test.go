package parser

import (
	"fmt"
	"reflect"
	"testing"
)

type tst struct {
	input string
	success bool
	output interface{}
}

func doTest(s tst, opts ...Option) func(t *testing.T) {
	return func(t *testing.T) {
		//t.Parallel()
		fmt.Printf("Testing \"%s\"\n", s.input)
		iface, err := Parse("test", []byte(s.input), opts...)
		if err != nil && s.success {
			fmt.Printf("Error: %s\n", err.Error())
			t.Fail()
			return
		} else if err == nil && !s.success {
			fmt.Printf("Expecting an error and got none: %#v\n", iface)
			t.Fail()
			return
		} else if err != nil && !s.success {
			fmt.Printf("Successfully got an error: %s\n", err.Error())
			return
		}
		typ := "nil"
		if iface != nil {
			typ = reflect.TypeOf(iface).String()
		}
		if reflect.TypeOf(iface) != reflect.TypeOf(s.output) {
			fmt.Printf("Type mismatch, expected %s but got %s\n", reflect.TypeOf(s.output), reflect.TypeOf(iface))
			t.Fail()
			return
		}
		if iface != s.output {
			fmt.Printf("Value mismatch expected %v but got %v\n", s.output, iface)
			t.Fail()
			return
		}
		fmt.Printf("Got %#v -> %s\n", iface, typ)
	}
}

func TestParse_Number(t *testing.T) {

	var tests = []tst{
		{"0", true, int64(0)},
		{"1", true, int64(1)},
		{"-1", true, int64(-1)},
		{"0.", true, float64(0)},
		{"1.", true, float64(1)},
		{"-1.", true, float64(-1)},
		{"0.0", false, nil},
		{"-0", false, nil},
		{"-0.0", false, nil},
		{"0.1", true, float64(0.1)},
		{"1.1", true, float64(1.1)},
		{"-1.1", true, float64(-1.1)},
		{"11", true, int64(11)},
		{"-11", true, int64(-11)},
		{"1234567890.0", true, float64(1234567890.0)},
		{"1234567890", true, int64(1234567890)},
	}

	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("Number")))
	}
}


func TestParse_String(t *testing.T) {
	var tests = []tst{
		{"''", true, ""},
		{"\"\"", true, ""},
		{"'hello'", true, "hello"},
		{"\"world\"", true, "world"},
		{"'π'", true, "π"},
		{"'it\\'s π day'", true, "it's π day" },
		{"\"it's π day\"", true, "it's π day"},
		{"'\"we\\'ll see about that \"'", true, "\"we'll see about that \""},
		{"\"\\\"we'll see about that \\\"\"", true, "\"we'll see about that \""},
		{"\"hello", false, nil},
		{"'hello", false, nil},
		{"hello\"", false, nil},
		{"hello'", false, nil},
		{"\"hello'", false, nil},
		{"hello", false, nil},
	}

	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("String")))
	}
}

func TestParse_UnquotedString(t *testing.T) {
	var tests = []tst{
		{"{{ hello }}", true, "{{ hello }}"},
		{"{{ hello-world }}", true, "{{ hello-world }}"},
		{"{{ h0la }}", true, "{{ h0la }}"},
		{"{{hello }}", true, "{{hello }}"},
		{"{{ hello}}", true, "{{ hello}}"},
		{"{{	hello	}}", true, "{{	hello	}}"},
		{"{{ hello", false, nil},
		{"{{ hello }", false, nil},
		{"{ hello }}", false, nil},
	}

	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("UnquotedString")))
	}
}

func TestParse_Value(t *testing.T) {
	var tests = []tst {
		{"1", true, Val{typ:ValTypeInt, itg:1}},
		{"1.0", true, Val{typ:ValTypeFloat, flt:1.0}},
		{"0", true, Val{typ:ValTypeInt, itg:0}},
		{"0.0", false, nil},
		{"'hello'", true, Val{typ:ValTypeString, str:"hello"}},
		{"hello", false, nil},
		{"true", true, Val{typ: ValTypeBool, bl: true}},
		{"false",true, Val{typ: ValTypeBool, bl: false}},
		{"nil", true, Val{typ: ValTypeNil}},
		{"null", true, Val{typ: ValTypeNil}},
		{"{{ .hello }}", true, Val{typ: ValTypeTemplate, str: "{{ .hello }}"}},
		{"'{{ .greet }}, World!'", true, Val{typ: ValTypeTemplate, str: "{{ .greet }}, World!"}},
		{"\"{'hello':{'some':'json','here':true}}\"", true, Val{typ: ValTypeString, str: "{'hello':{'some':'json','here':true}}"}},
		{"True", true, Val{typ: ValTypeBool, bl: true}},
		{"TRUE", true, Val{typ: ValTypeBool, bl: true}},
		{"False", true, Val{typ: ValTypeBool, bl: false}},
		{"FALSE", true, Val{typ: ValTypeBool, bl: false}},
		{"Nil", true, Val{typ: ValTypeNil}},
		{"Null", true, Val{typ: ValTypeNil}},
		{"NIL", true, Val{typ: ValTypeNil}},
		{"NULL", true, Val{typ: ValTypeNil}},
	}

	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("Value")))
	}
}

func TestParse_PredNumericOp_Equal(t *testing.T) {
	var arg = map[string]Val {
		"0":            {typ: ValTypeInt, itg: 0},
		"1":            {typ: ValTypeInt, itg: 1},
		"1.1":          {typ: ValTypeFloat, flt: 1.1},
		"2":            {typ: ValTypeInt, itg: 2},
		"2.0":          {typ: ValTypeFloat, flt: 2.0},
		"{{ .hello }}": {typ: ValTypeTemplate, str: "{{ .hello }}"},
	}
	var combiner = map[string]bool {
		"==": true, " ==": true, "== ": true, " == ": true, " equal ": true, " equals ": true, " Equal ": true, " Equals ": true, " EQUAL ": true, " EQUALS ": true, " EqUaLs ": true,
		"equal": false, "equal ": false, " equal": false, "equals": false,
		"	equal	": true, "     equal     ": true,
	}
	tests := make([]tst, 0)
	for i := range arg {
		for j := range combiner {
			for k := range arg{
				tests = append(tests, tst{fmt.Sprintf("%s%s%s", i, j, k), combiner[j], PredNumericOp{t: PredTypeNumberEqual, left: arg[i], right: arg[k]}})
			}
		}
	}
	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("Predicate")))
	}
}

func TestParse_PredNumericOp_Equal_WrongTypes(t *testing.T) {
	var arg = map[string]Val {
		"\"hello":   {typ: ValTypeString, str: "\"hello"},
		"\"1\"":   {typ: ValTypeString, str: "1"},
		"\"1.1\"": {typ: ValTypeString, str: "1.1"},
		"TRUE":    {},
		"False":   {},
		"Null":    {},
		"nil":     {},
	}
	var combiner = map[string]bool {
		"==": false, " ==": false, " equal ": false, " equals ": false, " Equal ": false, " Equals ": false, " EQUAL ": false,
	}
	tests := make([]tst, 0)
	for i := range arg {
		for j := range combiner {
			for k := range arg{
				tests = append(tests, tst{fmt.Sprintf("%s%s%s", i, j, k), false, nil})
			}
		}
	}
	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("Predicate")))
	}
}


func TestParse_PredNumericOp_NotEqual(t *testing.T) {
	var arg = map[string]Val {
		"0":            {typ: ValTypeInt, itg: 0},
		"1":            {typ: ValTypeInt, itg: 1},
		"1.1":          {typ: ValTypeFloat, flt: 1.1},
		"2":            {typ: ValTypeInt, itg: 2},
		"2.0":          {typ: ValTypeFloat, flt: 2.0},
		"{{ .hello }}": {typ: ValTypeTemplate, str: "{{ .hello }}"},
	}
	var combiner = map[string]bool {
		"!=": true, " !=": true, "!= ": true, " != ": true, " not equal ": true, " not equals ": true, " Not Equal ": true, " Not Equals ": true, " NOT EQUAL ": true, " NOT EQUALS ": true, " NoT EqUaLs ": true,
		"not equal": false, "notequal ": false, " not equal": false, "not equals": false, " notequal ": false,
		"	not equal	": true, "     not equal     ": true, "   not   equal   ": true,
	}
	tests := make([]tst, 0)
	for i := range arg {
		for j := range combiner {
			for k := range arg{
				tests = append(tests, tst{fmt.Sprintf("%s%s%s", i, j, k), combiner[j], PredNotOp{v: PredNumericOp{t: PredTypeNumberEqual, left: arg[i], right: arg[k]}}})
			}
		}
	}
	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("Predicate")))
	}
}

func TestParse_PredNumericOp_NotEqual_WrongTypes(t *testing.T) {
	var arg = map[string]Val {
		"\"hello":   {typ: ValTypeString, str: "\"hello"},
		"\"1\"":   {typ: ValTypeString, str: "1"},
		"\"1.1\"": {typ: ValTypeString, str: "1.1"},
		"TRUE":    {},
		"False":   {},
		"Null":    {},
		"nil":     {},
	}
	var combiner = map[string]bool {
		"!=": false, " !=": false, " not equal ": false, " not equals ": false, " Not Equal ": false, " Not Equals ": false, " NOT EQUAL ": false,
	}
	tests := make([]tst, 0)
	for i := range arg {
		for j := range combiner {
			for k := range arg{
				tests = append(tests, tst{fmt.Sprintf("%s%s%s", i, j, k), false, nil})
			}
		}
	}
	for _, s := range tests {
		t.Run(s.input, doTest(s, Entrypoint("Predicate")))
	}
}

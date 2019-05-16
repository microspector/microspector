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

func doTest(s tst, opts ...Option) (func(t *testing.T)) {
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
			fmt.Printf("Successfully got an error\n")
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
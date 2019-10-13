package parser

import (
	"fmt"
	"reflect"
	"strings"
)

type SetCommand struct {
	Name  string
	Value interface{}
	When  Expression
}

func (hc *SetCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	if strings.Contains(hc.Name, ".") {
		panic(fmt.Errorf("nested variables are not supported yet"))
	}

	if hc.Name == "State" {
		panic(fmt.Errorf(".State is a reserved variable for current state of the execution context"))
	}

	var i interface{}
	i = hc.Value

	for {
		t := reflect.TypeOf(i)
		if t.Implements(reflect.TypeOf((*Expression)(nil)).Elem()) {
			i = i.(Expression).Evaluate(l)
		} else {
			break
		}
		if i == nil {
			break
		}
	}
	l.GlobalVars[hc.Name] = i
	return i
}

func (hc *SetCommand) SetWhen(expr Expression) {
	hc.When = expr
}

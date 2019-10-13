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
	Async bool
}

func (sc *SetCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if sc.When != nil && !IsTrue(sc.When.Evaluate(l)) {
		return nil
	}

	if strings.Contains(sc.Name, ".") {
		panic(fmt.Errorf("nested variables are not supported yet"))
	}

	if sc.Name == "State" {
		panic(fmt.Errorf(".State is a reserved variable for current state of the execution context"))
	}

	var i interface{}
	i = sc.Value

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

	l.GlobalVars[sc.Name] = i
	return i
}

func (sc *SetCommand) SetWhen(expr Expression) {
	sc.When = expr
}

func (sc *SetCommand) SetAsync(async bool) {
	sc.Async = async
}

func (sc *SetCommand) IsAsync() bool {
	return sc.Async
}

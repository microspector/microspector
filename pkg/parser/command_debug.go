package parser

import (
	"fmt"
	"reflect"
)

type DebugCommand struct {
	Values interface{}
	When   Expression
	Async bool
}

func (dc *DebugCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if dc.When == nil || IsTrue(dc.When.Evaluate(l)) {
		t := reflect.TypeOf(dc.Values)
		if t.Implements(reflect.TypeOf((*Expression)(nil)).Elem()) {
			fmt.Printf("%+v\n", dc.Values.(Expression).Evaluate(l))
		} else {
			fmt.Printf("%+v\n", dc.Values)
		}
	}

	return nil
}

func (dc *DebugCommand) SetWhen(expr Expression) {
	dc.When = expr
}

func (dc *DebugCommand) SetAsync(async bool) {
	dc.Async = async
}

func (dc *DebugCommand) IsAsync() bool {
	return dc.Async
}


package parser

import (
	"fmt"
	"reflect"
)

type DebugCommand struct {
	Values ExprArray
	When   Expression
	Async  bool
}

func (dc *DebugCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if dc.When == nil || IsTrue(dc.When.Evaluate(l)) {
		for _, e := range dc.Values.Values {
			t := reflect.TypeOf(e)
			if t.Implements(reflect.TypeOf((*Expression)(nil)).Elem()) {
				fmt.Printf("%+v\n", e.(Expression).Evaluate(l))
			} else {
				fmt.Printf("%+v\n", e)
			}
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

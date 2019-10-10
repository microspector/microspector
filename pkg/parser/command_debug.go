package parser

import (
	"fmt"
	"reflect"
)

type DebugCommand struct {
	Values interface{}
	When   Expression
}

func (dc *DebugCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	//for _, y := range dc.Values {
	//	fmt.Printf("%+v ", y)
	//}
	t := reflect.TypeOf(dc.Values)
	if t.Implements(reflect.TypeOf((*Expression)(nil)).Elem()) {
		fmt.Printf("%+v\n", dc.Values.(Expression).Evaluate(l))
	} else {
		fmt.Printf("%+v\n", dc.Values)
	}

	return nil
}

func (dc *DebugCommand) SetWhen(expr Expression) {
	dc.When = expr
}

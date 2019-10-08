package parser

import "fmt"

type DebugCommand struct {
	Values []interface{}
	When   *Expression
}

func (dc *DebugCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	for _, y := range dc.Values {
		fmt.Printf("%+v ", y)
	}
	fmt.Printf("\n")
	return nil
}

func (dc *DebugCommand) SetWhen(expr *Expression) {
	dc.When = expr
}

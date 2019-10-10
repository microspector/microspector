package parser

import "fmt"

type EchoCommand struct {
	String string
	Values []interface{}
	When   Expression
}

func (ec *EchoCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	fmt.Printf(ec.String, ec.Values...)
	return nil
}

func (ec *EchoCommand) SetWhen(expr Expression) {
	ec.When = expr
}

package parser

import "fmt"

type EchoCommand struct {
	String string
	Values []interface{}
}

func (ec *EchoCommand) Run(l *lex) interface{} {
	defer l.wg.Done()
	fmt.Printf(ec.String, ec.Values...)
	return nil
}

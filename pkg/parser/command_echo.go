package parser

import "fmt"

type EchoCommand struct {
	Command
	String string
	Values []interface{}
	When   Expression
	Async  bool
}

func (ec *EchoCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	if ec.When == nil || IsTrue(ec.When.Evaluate(l)) {
		fmt.Printf(ec.String, ec.Values...)
	}
	return nil
}

func (ec *EchoCommand) SetWhen(expr Expression) {
	ec.When = expr
}

func (ec *EchoCommand) SetAsync(async bool) {
	ec.Async = async
}

func (ec *EchoCommand) IsAsync() bool {
	return ec.Async
}
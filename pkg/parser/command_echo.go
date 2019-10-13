package parser

import "fmt"

type EchoCommand struct {
	Command
	Expr   Expression
	Params ExprArray
	When   Expression
	Async  bool
}

func (ec *EchoCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	if ec.When == nil || IsTrue(ec.When.Evaluate(l)) {

		params := make([]interface{}, len(ec.Params.Values))
		for x, a := range ec.Params.Values {
			params[x] = a.Evaluate(l)
		}

		fmt.Printf(fmt.Sprintf("%v", ec.Expr.Evaluate(l)), params...)
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

package parser

import "errors"

var (
	ErrStopExecution = errors.New("unable to find the key")
)

type EndCommand struct {
	Expr  Expression
	When  Expression
	Async bool
}

func (ec *EndCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if ec.When == nil || IsTrue(ec.When.Evaluate(l)) {
		if ec.Expr == nil || IsTrue(ec.Expr.Evaluate(l)) {
			return ErrStopExecution
		}

	}
	return nil
}

func (ec *EndCommand) SetWhen(expr Expression) {
	ec.When = expr
}

func (ec *EndCommand) SetAsync(async bool) {
	ec.Async = async
}

func (ec *EndCommand) IsAsync() bool {
	return ec.Async
}

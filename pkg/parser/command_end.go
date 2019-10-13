package parser

import "errors"

var (
	ErrStopExecution = errors.New("unable to find the key")
)

type EndCommand struct {
	Expr Expression
	When Expression
}

func (ec *EndCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if ec.When == nil && ec.Expr == nil {
		return ErrStopExecution
	}

	if ec.Expr == nil && IsTrue(ec.When.Evaluate(l)) {
		return ErrStopExecution
	}

	if ec.When == nil && IsTrue(ec.Expr.Evaluate(l)) {
		return ErrStopExecution
	}

	return nil
}

func (ec *EndCommand) SetWhen(expr Expression) {
	ec.When = expr
}

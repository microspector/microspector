package parser

type AssertCommand struct {
	Command
	Expr   Expression
	When   Expression
}

func (ac *AssertCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	r := IsTrue(ac.Expr.Evaluate(l))

	if r {
		l.State.Assert.Success++
	} else {
		l.State.Assert.Fail++
	}
	return r
}

func (ac *AssertCommand) SetWhen(expr Expression) {
	ac.When = expr
}

package parser

type ShouldCommand struct {
	Command
	Failed bool
	Expr   Expression
	When   Expression
}

func (sc *ShouldCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	r := false

	if sc.When == nil || IsTrue(sc.When.Evaluate(l)) {
		r = IsTrue(sc.Expr.Evaluate(l))
		if r {
			l.State.Should.Success++
		} else {
			l.State.Should.Fail++
		}
	}

	return r
}

func (sc *ShouldCommand) SetWhen(expr Expression) {
	sc.When = expr
}

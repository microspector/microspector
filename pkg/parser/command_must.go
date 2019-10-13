package parser

type MustCommand struct {
	Failed bool
	Expr   Expression
	When   Expression
}

func (mc *MustCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	r := false

	if mc.When == nil || IsTrue(mc.When.Evaluate(l)) {
		r = IsTrue(mc.Expr.Evaluate(l))

		if r {
			l.State.Must.Success++
		} else {
			l.State.Must.Fail++
		}
	}

	return r
}

func (mc *MustCommand) SetWhen(expr Expression) {
	mc.When = expr
}

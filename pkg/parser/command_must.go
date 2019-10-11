package parser

type MustCommand struct {
	Failed bool
	Expr   Expression
	When   Expression
}

func (mc *MustCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	r := IsTrue(mc.Expr.Evaluate(l))

	if r {
		l.State.Must.Success++
	} else {
		l.State.Must.Fail++
	}

	return r
}

func (mc *MustCommand) SetWhen(expr Expression) {
	mc.When = expr
}

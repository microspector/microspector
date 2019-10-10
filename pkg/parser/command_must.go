package parser

type MustCommand struct {
	Failed bool
	When   Expression
}

func (mc *MustCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got a MUST Command here"
}

func (mc *MustCommand) SetWhen(expr Expression) {
	mc.When = expr
}

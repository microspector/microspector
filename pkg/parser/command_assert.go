package parser

type AssertCommand struct {
	Failed bool
	When   *Expression
}

func (ac *AssertCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got an ASSERT Command here"
}

func (ac *AssertCommand) SetWhen(expr *Expression) {
	ac.When = expr
}

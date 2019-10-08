package parser

type ShouldCommand struct {
	Command
	Failed bool
	When   *Expression
}

func (sc *ShouldCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got a SHOULD Command here"
}

func (sc *ShouldCommand) SetWhen(expr *Expression) {
	sc.When = expr
}

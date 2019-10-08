package parser

type EndCommand struct {
	When *Expression
}

func (ec *EndCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got a SET Command here"
}

func (ec *EndCommand) SetWhen(expr *Expression) {
	ec.When = expr
}

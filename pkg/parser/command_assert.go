package parser

type AssertCommand struct {
}

func (hc *AssertCommand) Run(l *lex) interface{} {
	return "we got an ASSERT Command here"
}

package parser

type AssertCommand struct {
	Failed bool
}

func (hc *AssertCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got an ASSERT Command here"
}

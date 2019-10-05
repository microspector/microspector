package parser

type AssertCommand struct {
	Failed bool
}

func (hc *AssertCommand) Run(l *lex) interface{} {
	defer l.wg.Done()
	return "we got an ASSERT Command here"
}

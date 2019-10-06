package parser

type EndCommand struct {
}

func (hc *EndCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got a SET Command here"
}

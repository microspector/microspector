package parser

type EndCommand struct {
}

func (hc *EndCommand) Run(l *lex) interface{} {
	defer l.wg.Done()
	return "we got a SET Command here"
}

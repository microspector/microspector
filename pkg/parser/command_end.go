package parser

type EndCommand struct {
}

func (hc *EndCommand) Run(l *lex) interface{} {
	return "we got a SET Command here"
}

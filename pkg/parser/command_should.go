package parser

type ShouldCommand struct {
}

func (hc *ShouldCommand) Run(l *lex) interface{} {
	return "we got a SHOULD Command here"
}

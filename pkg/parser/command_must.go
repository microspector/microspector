package parser

type MustCommand struct {
}

func (hc *MustCommand) Run(l *lex) interface{} {
	return "we got a MUST Command here"
}

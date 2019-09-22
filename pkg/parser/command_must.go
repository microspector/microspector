package parser

type MustCommand struct {
	Failed bool
}

func (hc *MustCommand) Run(l *lex) interface{} {
	return "we got a MUST Command here"
}
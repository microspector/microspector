package parser

type ShouldCommand struct {
	Command
	Failed bool
}

func (sc *ShouldCommand) Run(l *lex) interface{} {
	return "we got a SHOULD Command here"
}

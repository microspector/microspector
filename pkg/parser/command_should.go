package parser

type ShouldCommand struct {
	Command
	Failed bool
}

func (sc *ShouldCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	return "we got a SHOULD Command here"
}

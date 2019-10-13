package parser

type ExprFloat struct {
	Val float64
}

func (ef *ExprFloat) Evaluate(lexer *Lexer) interface{} {
	return ef.Val
}

package parser

type ExprFloat struct {
	Val float64
}

func (es *ExprFloat) Evaluate(lexer *Lexer) interface{} {
	return es.Val
}
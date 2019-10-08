package parser

type ExprString struct {
	Val string
}

func (es *ExprString) Evaluate(lexer *Lexer) interface{} {
	return es.Val
}

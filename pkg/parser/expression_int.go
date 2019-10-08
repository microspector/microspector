package parser

type ExprInteger struct {
	Val int64
}

func (es *ExprInteger) Evaluate(lexer *Lexer) interface{} {
	return es.Val
}

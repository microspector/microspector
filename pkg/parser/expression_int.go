package parser

type ExprInteger struct {
	Val int64
}

func (ei *ExprInteger) Evaluate(lexer *Lexer) interface{} {
	return ei.Val
}

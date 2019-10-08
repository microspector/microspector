package parser

type ExprBool struct {
	Val bool
}

func (b *ExprBool) Evaluate(lexer *Lexer) interface{} {
	return b.Val
}

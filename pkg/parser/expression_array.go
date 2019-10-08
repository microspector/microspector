package parser

type ExprArray struct {
	Values []*Expression
}

func (a *ExprArray) Evaluate(lexer *Lexer) interface{} {
	return a.Values
}
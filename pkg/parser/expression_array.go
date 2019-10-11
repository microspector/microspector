package parser

type ExprArray struct {
	Values []Expression
}

func (a *ExprArray) Evaluate(lexer *Lexer) interface{} {
	array := make([]interface{}, len(a.Values))
	for x, a := range a.Values {
		array[x] = a.Evaluate(lexer)
	}
	return array
}

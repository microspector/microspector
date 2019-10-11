package parser

type ExprFunc struct {
	Name   string
	Params []Expression
}

func (f *ExprFunc) Evaluate(lexer *Lexer) interface{} {
	array := make([]interface{}, len(f.Params))
	for x, a := range f.Params {
		array[x] = a.Evaluate(lexer)
	}
	return funcCall(f.Name, array)
}

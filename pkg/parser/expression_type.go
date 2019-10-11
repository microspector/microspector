package parser

type ExprType struct {
	Name string
}

func (et *ExprType) Evaluate(lexer *Lexer) interface{} {
	return et.Name
}

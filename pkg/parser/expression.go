package parser

type Expression interface {
	Evaluate(lexer *Lexer) interface{}
}

package parser

type Command interface {
	Run(l *lex) interface{}
}
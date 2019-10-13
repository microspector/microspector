package parser

type Command interface {
	Run(l *Lexer) interface{}
	SetWhen(exp Expression)
}

type IntoCommand interface {
	SetInto(into string)
}
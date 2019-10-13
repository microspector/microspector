package parser

type Command interface {
	Run(l *Lexer) interface{}
	SetWhen(exp Expression)
	SetAsync(async bool)
	IsAsync() bool
}

type IntoCommand interface {
	SetInto(into string)
}

type AssertionCommand interface {
	SetAssertionMessage(expression Expression)
}

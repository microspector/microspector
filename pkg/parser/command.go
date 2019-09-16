package parser

type Command interface {
	Run() interface{}
}

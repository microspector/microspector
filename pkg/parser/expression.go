package parser

type Expression interface {
	Evaluate() interface{}
}
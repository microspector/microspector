package parser

import (
	"github.com/microspector/microspector/pkg/lookup"
)

type ExprVariable struct {
	Name string
}

func (v *ExprVariable) Evaluate(lexer *Lexer) interface{} {
	l, err := lookup.Lookup(lexer.GlobalVars, v.Name)
	if err != nil {
		return nil
	}
	return l.Interface()
}
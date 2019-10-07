package parser

import "github.com/microspector/microspector/pkg/lookup"

type Variable struct {
	Name string
}

func (v *Variable) Evaluate(lexer *Lexer) interface{} {
	l, err := lookup.Lookup(lexer.GlobalVars, v.Name)
	if err != nil {
		return nil
	}
	return l.Interface()
}

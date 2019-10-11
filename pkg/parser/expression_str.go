package parser

import "github.com/microspector/microspector/pkg/templating"

type ExprString struct {
	Val string
}

func (es *ExprString) Evaluate(lexer *Lexer) interface{} {

	if isTemplate(es.Val) {
		str, err := templating.ExecuteTemplate(es.Val, lexer.GlobalVars)
		if err != nil {
			///TODO: error message to print out?
		}

		return str
	}

	return es.Val
}

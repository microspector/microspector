package parser

import (
	"github.com/microspector/microspector/pkg/lookup"
	"strings"
)

type ExprVariable struct {
	Name string
}

func (ev *ExprVariable) Evaluate(lexer *Lexer) interface{} {

	// we do not need to run that heavy reflect func if we don't have punctuation in var name
	if !strings.Contains(ev.Name, ".") {
		ok, x := lexer.GlobalVars[ev.Name]
		if x {
			return ok
		}
		return nil
	}

	l, err := lookup.LookupString(lexer.GlobalVars, ev.Name)

	if err != nil || !l.IsValid() {
		return nil
	}

	return l.Interface()
}

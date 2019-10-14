package parser

import (
	"fmt"
	"io/ioutil"
	"os"
)

type IncludeCommand struct {
	Expr  Expression
	When  Expression
	Async bool
}

func (ic *IncludeCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if ic.When != nil && !IsTrue(ic.When.Evaluate(l)) {
		return nil
	}

	bytes, err := ioutil.ReadFile(ic.Expr.Evaluate(l).(string))

	if err != nil {
		fmt.Println(fmt.Errorf("error including file: %s", err))
		os.Exit(1)
	}
	//todo: make only 1 pointer that holds state and globalvars to pass here.
	lex := Parse(string(bytes))
	lex.State = l.State
	lex.GlobalVars = l.GlobalVars
	Run(lex)
	l.State = lex.State
	l.GlobalVars = lex.GlobalVars

	return nil
}

func (ic *IncludeCommand) IsAsync() bool {
	return ic.Async
}

func (ic *IncludeCommand) SetAsync(async bool) {
	ic.Async = async
}

func (ic *IncludeCommand) SetWhen(expr Expression) {
	ic.When = expr
}

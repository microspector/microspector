package parser

import (
	"fmt"
	"io/ioutil"
	"os"
)

type IncludeCommand struct {
	File  string
	When  Expression
	Async bool
}

func (ic *IncludeCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if ic.When == nil || IsTrue(ic.When.Evaluate(l)) {

		bytes, err := ioutil.ReadFile(ic.File)

		if err != nil {
			fmt.Println(fmt.Errorf("error including file: %s", err))
			os.Exit(1)
		}
		lex := Parse(string(bytes))
		lex.State = l.State
		lex.GlobalVars = l.GlobalVars
		Run(lex)
	}

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

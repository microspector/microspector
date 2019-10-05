package parser

import (
	"fmt"
	"io/ioutil"
	"os"
)

type IncludeCommand struct {
	File string
}

func (ic *IncludeCommand) Run(l *lex) interface{} {
	defer l.wg.Done()

	bytes, err := ioutil.ReadFile(ic.File)

	if err != nil {
		fmt.Println(fmt.Errorf("error including file: %s", err))
		os.Exit(1)
	}
	lex := Parse(string(bytes))
	lex.State = l.State
	lex.GlobalVars = l.GlobalVars
	Run(lex)

	return nil

}

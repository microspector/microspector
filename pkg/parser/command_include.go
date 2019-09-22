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

	bytes, err := ioutil.ReadFile(ic.File)

	if err != nil {
		fmt.Println(fmt.Errorf("error including file: %s", err))
		os.Exit(1)
	}
	lex := Parse(string(bytes))
	Run(lex)

	return nil

}

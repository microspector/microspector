package parser

import (
	"fmt"
	"strings"
)

type SetCommand struct {
	Name  string
	Value interface{}
}

func (hc *SetCommand) Run(l *lex) interface{} {
	if strings.Contains(hc.Name, ".") {
		panic(fmt.Errorf("nested variables are not supported yet"))
	}

	if hc.Name == "State" {
		panic(fmt.Errorf(".State is a reserved variable for current state of the execution context"))
	}

	l.GlobalVars[hc.Name] = hc.Value
	return hc.Value
}
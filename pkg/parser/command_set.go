package parser

import (
	"fmt"
	"strings"
)

type SetCommand struct {
	Name  string
	Value interface{}
	When  *Expression
}

func (hc *SetCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	if strings.Contains(hc.Name, ".") {
		panic(fmt.Errorf("nested variables are not supported yet"))
	}

	if hc.Name == "State" {
		panic(fmt.Errorf(".State is a reserved variable for current state of the execution context"))
	}

	l.GlobalVars[hc.Name] = hc.Value
	return hc.Value
}

func (hc *SetCommand) SetWhen(expr *Expression) {
	hc.When = expr
}

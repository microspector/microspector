package parser

import (
	"fmt"
	"strings"
)

type SetCommand struct {
	Name  string
	Value interface{}
}

func (hc *SetCommand) Run() interface{} {
	if strings.Contains(hc.Name, ".") {
		panic(fmt.Errorf("nested variables are not supported yet"))
	}

	GlobalVars[hc.Name] = hc.Value
	return hc.Value
}

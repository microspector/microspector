package parser

import (
	"fmt"
	"os/exec"
	"strings"
)

type CmdCommand struct {
	Params []interface{}
}

func (cc *CmdCommand) Run(l *lex) interface{} {

	var params []string

	for _, p := range cc.Params {
		params = append(params, fmt.Sprintf("%v", p))
	}

	var cm *exec.Cmd

	if len(cc.Params) > 1 {
		cm = exec.Command(params[0], params[1:]...)
	} else if len(cc.Params) == 1 {
		cm = exec.Command(params[0])
	}

	out, err := cm.Output()

	if err != nil {
		return err
	}

	return strings.TrimSpace(string(out))
}

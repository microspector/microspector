package parser

import (
	"fmt"
	"os/exec"
	"strings"
)

type CmdCommand struct {
	Command
	Params ExprArray
	When   Expression
	Into   string
}

func (cc *CmdCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if cc.When == nil || IsTrue(cc.When.Evaluate(l)) {

		params := make([]string, len(cc.Params.Values))
		for x, a := range cc.Params.Values {
			params[x] = fmt.Sprintf("%v", a.Evaluate(l))
		}

		var cm *exec.Cmd

		if len(cc.Params.Values) > 1 {
			cm = exec.Command(params[0], params[1:]...)
		} else if len(cc.Params.Values) == 1 {
			cm = exec.Command(params[0])
		}

		out, err := cm.Output()

		if err != nil {
			return err
		}
		l.GlobalVars[cc.Into] = strings.TrimSpace(string(out))
		return strings.TrimSpace(string(out))
	}

	return nil
}

func (cc *CmdCommand) SetWhen(expr Expression) {
	cc.When = expr
}

func (cc *CmdCommand) SetInto(into string) {
	cc.Into = into
}

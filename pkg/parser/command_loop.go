package parser

type LoopCommand struct {
	Command
	Commands []Command
	When     Expression
	Async    bool
	Var      ExprVariable
	In       ExprVariable
}

func (lc *LoopCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	if lc.When != nil && !IsTrue(lc.When.Evaluate(l)) {
		return nil
	}

	rng := lc.In.Evaluate(l)
	switch rng.(type) {
	case []interface{}:
		for _, val := range rng.([]interface{}) {
			l.GlobalVars[lc.Var.Name] = val
			for _, cm := range lc.Commands {
				l.wg.Add(1)
				if cm.IsAsync() {
					go cm.Run(l)
				} else {
					r := cm.Run(l)
					if r == ErrStopExecution {
						return ErrStopExecution
					}
				}
			}
		}

		break

	case []string:
		for _, val := range rng.([]string) {
			l.GlobalVars[lc.Var.Name] = val
			for _, cm := range lc.Commands {
				l.wg.Add(1)
				if cm.IsAsync() {
					go cm.Run(l)
				} else {
					r := cm.Run(l)
					if r == ErrStopExecution {
						return ErrStopExecution
					}
				}
			}
		}
		break

	case []int:
		for _, val := range rng.([]int) {
			l.GlobalVars[lc.Var.Name] = val
			for _, cm := range lc.Commands {
				l.wg.Add(1)
				if cm.IsAsync() {
					go cm.Run(l)
				} else {
					r := cm.Run(l)
					if r == ErrStopExecution {
						return ErrStopExecution
					}
				}
			}
		}
		break

	case map[string]interface{}:
		for _, val := range rng.(map[string]interface{}) {
			l.GlobalVars[lc.Var.Name] = val
			for _, cm := range lc.Commands {
				l.wg.Add(1)
				if cm.IsAsync() {
					go cm.Run(l)
				} else {
					r := cm.Run(l)
					if r == ErrStopExecution {
						return ErrStopExecution
					}
				}
			}
		}

		break
	}

	return nil
}

func (lc *LoopCommand) SetWhen(expr Expression) {
	lc.When = expr
}

func (lc *LoopCommand) SetAsync(async bool) {
	lc.Async = async
}

func (lc *LoopCommand) IsAsync() bool {
	return lc.Async
}

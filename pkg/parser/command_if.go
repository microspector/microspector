package parser

type IfCommand struct {
	Command
	IfCommands   []Command
	When         Expression
	Async        bool
	ElseCommands []Command
	Predicate    Expression
}

func (ic *IfCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	if ic.When != nil && !IsTrue(ic.When.Evaluate(l)) {
		return nil
	}
	if IsTrue(ic.Predicate.Evaluate(l)) {

		for _, cm := range ic.IfCommands {
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

	} else {
		if ic.ElseCommands != nil {
			for _, cm := range ic.ElseCommands {
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
	}

	return nil
}

func (ic *IfCommand) SetWhen(expr Expression) {
	ic.When = expr
}

func (ic *IfCommand) SetAsync(async bool) {
	ic.Async = async
}

func (ic *IfCommand) IsAsync() bool {
	return ic.Async
}

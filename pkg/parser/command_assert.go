package parser

type AssertCommand struct {
	Command
	Expr    Expression
	When    Expression
	Async   bool
	Message Expression
}

func (ac *AssertCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	r := false

	if ac.When == nil || IsTrue(ac.When.Evaluate(l)) {
		r = IsTrue(ac.Expr.Evaluate(l))

		if r {
			l.State.Assert.Success++
		} else {

			if ac.Message != nil {
				l.State.Assert.Messages = append(l.State.Assert.Messages, ac.Message.Evaluate(l).(string))
			}

			l.State.Assert.Fail++
		}
	}
	return r
}

func (ac *AssertCommand) SetWhen(expr Expression) {
	ac.When = expr
}

func (ac *AssertCommand) SetAsync(async bool) {
	ac.Async = async
}

func (ac *AssertCommand) IsAsync() bool {
	return ac.Async
}

func (ac *AssertCommand) SetAssertionMessage(expression Expression) {
	ac.Message = expression
}

package parser

type ShouldCommand struct {
	Command
	Expr    Expression
	When    Expression
	Async   bool
	Message Expression
}

func (sc *ShouldCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	r := false

	if sc.When == nil || IsTrue(sc.When.Evaluate(l)) {
		r = IsTrue(sc.Expr.Evaluate(l))
		if r {
			l.State.Should.Success++
		} else {
			if sc.Message != nil {
				l.State.Should.Messages = append(l.State.Should.Messages, sc.Message.Evaluate(l).(string))
			}
			l.State.Should.Fail++
		}
	}

	return r
}

func (sc *ShouldCommand) SetWhen(expr Expression) {
	sc.When = expr
}

func (sc *ShouldCommand) SetAsync(async bool) {
	sc.Async = async
}

func (sc *ShouldCommand) IsAsync() bool {
	return sc.Async
}

func (sc *ShouldCommand) SetAssertionMessage(expression Expression) {
	sc.Message = expression
}

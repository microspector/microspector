package parser

type MustCommand struct {
	Expr    Expression
	When    Expression
	Async   bool
	Message Expression
}

func (mc *MustCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()

	r := false

	if mc.When == nil || IsTrue(mc.When.Evaluate(l)) {
		r = IsTrue(mc.Expr.Evaluate(l))

		if r {
			l.State.Must.Success++
		} else {
			if mc.Message != nil {
				l.State.Must.Messages = append(l.State.Must.Messages, mc.Message.Evaluate(l).(string))
			}
			l.State.Must.Fail++
		}
	}

	return r
}

func (mc *MustCommand) IsAsync() bool {
	return mc.Async
}

func (mc *MustCommand) SetWhen(expr Expression) {
	mc.When = expr
}

func (mc *MustCommand) SetAsync(async bool) {
	mc.Async = async
}

func (mc *MustCommand) SetAssertionMessage(expression Expression) {
	mc.Message = expression
}

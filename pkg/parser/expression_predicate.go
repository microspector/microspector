package parser

type ExprPredicate struct {
	Left     Expression
	Operator string // Token
	Right    Expression
	Not      bool
}

func (ep *ExprPredicate) Evaluate(lexer *Lexer) interface{} {
	r := RunOp(ep.Left.Evaluate(lexer), ep.Operator, ep.Right.Evaluate(lexer))
	if ep.Not {
		return !r
	}
	return r
}

package parser

type ExprPredicate struct {
	Left     Expression
	Operator string // Token
	Right    Expression
	Not      bool
}

func (p *ExprPredicate) Evaluate(lexer *Lexer) interface{} {
	r := runop(p.Left.Evaluate(lexer), p.Operator, p.Right.Evaluate(lexer))
	if p.Not {
		return !r
	}
	return r
}

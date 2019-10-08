package parser

type ExprPredicate struct {
	Left     Expression
	Operator string // Token
	Right    Expression
}

func (p *ExprPredicate) Evaluate(lexer *Lexer) interface{} {
	return runOpPositive(p.Left.Evaluate(lexer),p.Operator,p.Right.Evaluate(lexer))
}

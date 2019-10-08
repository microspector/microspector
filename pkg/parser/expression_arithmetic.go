package parser

import (
	"fmt"
	"log"
)

type ExprArithmetic struct {
	Left     Expression
	Operator string // Token
	Right    Expression
}

func (a *ExprArithmetic) Evaluate(lexer *Lexer) interface{} {
	switch a.Operator {
	case "+":
		return add(a.Right.Evaluate(lexer), a.Left.Evaluate(lexer))
	case "-":
		return subtract(a.Right.Evaluate(lexer), a.Left.Evaluate(lexer))
	case "*":
		return multiply(a.Right.Evaluate(lexer), a.Left.Evaluate(lexer))
	case "/":
		return divide(a.Right.Evaluate(lexer), a.Left.Evaluate(lexer))
	default:
		log.Fatal(fmt.Errorf("unknown operator: %s", a.Operator))
		return nil
	}
}

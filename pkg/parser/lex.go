package parser

import "log"

type lex struct {
	tokens     chan Token
	GlobalVars map[string]interface{}
	State      Stats
}

func (l *lex) All() []Token {
	tokens := make([]Token, 0)

	for {

		v := <-l.tokens
		if v.Type == EOF || v.Type == -1 {
			break
		}

		tokens = append(tokens, v)
	}

	return tokens
}

func (l *lex) Lex(lval *yySymType) int {
	v := <-l.tokens
	if v.Type == EOF || v.Type == -1 {
		return 0
	}
	lval.val = v.Val
	return v.Type
}

func (l *lex) Error(e string) {
	log.Fatal(e)
}

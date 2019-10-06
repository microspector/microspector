package parser

import (
	"log"
	"sync"
)

type Lexer struct {
	tokens     chan Token
	GlobalVars map[string]interface{}
	State      Stats
	wg         *sync.WaitGroup
}

func (l *Lexer) All() []Token {
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

func (l *Lexer) Lex(lval *yySymType) int {
	v := <-l.tokens
	if v.Type == EOF || v.Type == -1 {
		return 0
	}
	lval.val = v.Val
	return v.Type
}

func (l *Lexer) Error(e string) {
	log.Println(e)
}

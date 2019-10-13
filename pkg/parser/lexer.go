package parser

import (
	"fmt"
	"log"
	"sync"
)

type Token struct {
	Type int
	Val  interface{}
}

func (t *Token) TypeName() string {
	if name, ok := keywordStrings[t.Type]; ok {
		return name
	}

	return fmt.Sprintf("%s", t.Val)
}

func (t *Token) Error() string {
	panic("implement me")
}

func (t *Token) String() string {
	return fmt.Sprintf("Token: %s, Val:%s", t.TypeName(), t.Val)
}

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

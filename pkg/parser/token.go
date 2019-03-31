package parser

import (
	"strings"
)

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	EOL
	KEYWORD
	STRING
	INTEGER
	OPERATOR
	ARRAY
	VARIABLE
	VARSTART
	VAREND
	WHITESPACE
	COMMAND
)

type Token struct {
	Position int
	Type     TokenType
	Text     string
	Tree     []Token
}

var constLookup = map[TokenType]string{
	ILLEGAL:    `Illegal Token`,
	EOF:        `End of File`,
	KEYWORD:    `KEYWORD`,
	STRING:     `STRING`,
	INTEGER:    `INTEGER`,
	OPERATOR:   `OPERATOR`,
	ARRAY:      `ARRAY`,
	VARIABLE:   `VARIABLE`,
	VARSTART:   `VARSTART`,
	VAREND:     `VAREND`,
	WHITESPACE: `WHITESPACE`,
	COMMAND:    `COMMAND`,
}

func (t *Token) TypeName() string {
	return constLookup[t.Type]
}

func (t *Token) Error() string {
	panic("implement me")
}

func (t *Token) IsCommand() bool {
	return t.Type == COMMAND
}

func (t *Token) IsWhiteSpace() bool {
	return t.Type == WHITESPACE
}

func (t *Token) IsIllegal() bool {
	return t.Type == ILLEGAL
}

func (t *Token) IsLegal() bool {
	return !t.IsIllegal()
}

func (t *Token) IsKeyword() bool {
	return t.Type == KEYWORD
}

func (t *Token) isQuotedString() bool {
	return t.Type == STRING
}

func (t *Token) isTemplated() bool {
	return t.isQuotedString() && strings.Contains(t.Text, "{{") && strings.Contains(t.Text, "}}")
}

func (t *Token) Tokenize() {
	s := NewScanner(strings.NewReader(t.Text))
	for {
		token := s.Scan()
		if token.Type == EOF || token.Type == ILLEGAL {
			break
		}
		t.Tree = append(t.Tree, token)
	}
}

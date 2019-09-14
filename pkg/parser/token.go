package parser

import "fmt"

type Token struct {
	Type int
	Text string
}

type Tokens []Token

func (t *Tokens) String() string {
	return fmt.Sprintf("no string bro %s", "x")
}

func (t *Token) TypeName() string {
	if name, ok := keywordStrings[t.Type]; ok {
		return name
	}

	return t.Text
}

func (t *Token) Error() string {
	panic("implement me")
}

func (t *Token) String() string {
	return fmt.Sprintf("Token: %s, Text:%s", t.TypeName(), t.Text)
}

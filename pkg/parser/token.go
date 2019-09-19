package parser

import "fmt"

type Token struct {
	Type int
	Val  interface{}
}

type Tokens []Token

func (t *Tokens) String() string {
	return fmt.Sprintf("no string bro %s", "x")
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

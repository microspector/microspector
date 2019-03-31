package parser

import (
	"fmt"
	"log"
)

type Set struct {
	Token Token
}

func (s *Set) Run(state State) error {
	argCount := 2

	if len(s.Token.Tree)-1 != argCount {
		log.Printf("SET command gets %d arguments %d given\n", argCount, len(s.Token.Tree)-1)
		return fmt.Errorf("SET command gets %d arguments %d given", argCount, len(s.Token.Tree)-1)
	}

	variableNameToken := s.Token.Tree[1]
	variableValueToken := s.Token.Tree[2]

	switch variableValueToken.Type {
	case VARIABLE:
		state.Vars[variableNameToken.Text] = query(variableValueToken.Text, state.Vars)
		break
	case STRING:
		if variableValueToken.isTemplated() {
			state.Vars[variableNameToken.Text], _ = executeTemplate(variableValueToken.Text, state.Vars)
		} else {
			state.Vars[variableNameToken.Text] = variableValueToken.Text
		}
		break
	}

	log.Printf("SET variable %s to %s\n", variableNameToken.Text, variableValueToken.Text)

	return nil
}

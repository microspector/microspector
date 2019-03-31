package parser

import (
	"github.com/davecgh/go-spew/spew"
	"log"
)

type Debug struct {
	Token Token
}

func (s *Debug) Run(state *State) error {

	for i := 0; i < len(s.Token.Tree); i++ {
		token := s.Token.Tree[i]

		switch token.Type {
		case STRING:
			if token.isTemplated() {

				_evaluated, err := executeTemplate(token.Text, state.Vars)

				if err != nil {
					log.Println(err)
				}

				log.Printf("[DEBUG] %s\n", _evaluated)

			} else {
				log.Printf("[DEBUG] %s\n", token.Text)
			}
			break

		case VARIABLE:
			//find value of VARIABLE from state
			log.Printf("[DEBUG] Value of : %s\n", token.Text)
			spew.Dump( query(token.Text, state.Vars) )
			break

		default:
			log.Printf("[DEBUG] %s\n", token.Text)
			break

		}
	}

	return nil
}

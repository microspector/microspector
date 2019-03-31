package parser

import "log"

type Microspector struct {
	Token Token
}

func (m *Microspector) Run(state *State) error {
	log.Println("MICROSPECTOR COMMAND")
	return nil
}

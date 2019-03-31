package main

import (
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
	"log"
)

func main() {

	var state = parser.State{
		Vars: map[string]interface{}{},
	}

	p := parser.NewParser("../tasks/main.msf")
	for _, cmd := range p.Commands {
		err := cmd.Run(state)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// print a report here.
	//spew.Dump(state)
}

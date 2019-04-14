package main

import (
	"flag"
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
	"log"
)

func main() {

	var state = &parser.State{
		Vars: map[string]interface{}{},
	}

	var file = flag.String("file", "../tasks/main.msf", "Task file path")
	flag.Parse()

	p := parser.NewParser(*file)
	for _, cmd := range p.Commands {
		err := cmd.Run(state)
		if err != nil {
			log.Println(err)
		}
	}

	// print a report here.

	log.Printf("Passed MUST : %d", state.SuccessMust)
	log.Printf("Passed SHOULD : %d", state.SuccessShould)
	log.Printf("Failed SHOULD : %d", state.FailedShould)
}

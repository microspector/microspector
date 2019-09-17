package main

import (
	"flag"
	"fmt"
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
	"io/ioutil"
	"os"
)

func main() {

	var file = flag.String("file", "../tasks/main.msf", "Task file path")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(fmt.Errorf("error reading file: %s", err))
		os.Exit(1)
	}

	parser.Run(parser.Parse(string(bytes)))

	//fmt.Printf("%+v\n", parser.GlobalVars)

}

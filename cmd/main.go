package main

import (
	"flag"
	"fmt"
	"github.com/microspector/microspector/pkg/parser"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	Version = "development"
	Build   = "master-dev"
)

func main() {

	parser.Version = Version
	parser.Build = Build

	var file, folder string

	var fi = flag.String("file", "", "task file path")
	var fo = flag.String("folder", "", "tasks folder path")
	var v = flag.Bool("version", false, "prints version")
	var vv = flag.Bool("verbose", false, "print out logs")
	flag.Parse()

	file = *fi
	folder = *fo
	parser.Verbose = *vv

	if *v {
		fmt.Printf("Microspector v%s (%s)\n", Version, Build)
		os.Exit(0)
	}

	if file == "" && folder == "" {
		flag.PrintDefaults()
	}

	var files = make([]string, 0)

	if folder != "" {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			fmt.Println(file)
		}
	} else if file != "" {
		files = append(files, file)
	}

	for _, f := range files {

		bytes, err := ioutil.ReadFile(f)

		if err != nil {
			fmt.Println(fmt.Errorf("error reading file: %s", err))
			os.Exit(1)
		}
		lex := parser.Parse(string(bytes))
		if *vv {
			fmt.Printf("%+v\n", lex)
		}
		parser.Run(lex)

		if *vv {
			fmt.Printf("%+v\n", lex.State)
			fmt.Printf("%+v\n", lex.GlobalVars)
		}
	}
}

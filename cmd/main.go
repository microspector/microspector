package main

import (
	"flag"
	"fmt"
	"github.com/tufanbarisyildirim/microspector/pkg/parser"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	Version = "development"
	Build   = "master-dev"
)

func main() {

	parser.Version = Version
	parser.Build = Build

	var file, folder string
	var err error

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

	files := make([]os.FileInfo, 0)

	if folder != "" {
		files, err = ioutil.ReadDir(folder)
		if err != nil {
			log.Fatal(err)
		}
	} else if file != "" {
		files = make([]os.FileInfo, 1)
		f, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		files[0] = f
	}

	for _, f := range files {

		if f.IsDir() {
			continue
		}

		bytes, err := ioutil.ReadFile(path.Join(folder, f.Name()))

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

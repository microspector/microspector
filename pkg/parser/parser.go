package parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Parser struct {
	s        *Scanner
	Commands []Command
}

func NewParser(file string) *Parser {

	p := &Parser{
		Commands: []Command{},
	}

	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)

	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		t, err := makeCommand(line)
		log.Println(t.Text)
		if err != nil {
			log.Fatalf("error on line x: %v", err)
		}

		if len(t.Tree) > 0 {

			switch strings.ToUpper(t.Tree[0].Text) {
			case "HTTP":
				p.Commands = append(p.Commands, &Http{
					Token: t,
				})
				break

			case "MUST", "SHOULD", "MUSTNOT", "SHOULDNOT":
				p.Commands = append(p.Commands, &Must{
					Token: t,
				})
				break

			case "SET":
				p.Commands = append(p.Commands, &Set{
					Token: t,
				})
				break

			case "MICROSPECTOR":
			case "ENDMICROSPECTOR":
				p.Commands = append(p.Commands, &Microspector{
					Token: t,
				})
				break

			default:
				log.Fatalf("Unknown Command %v\n", t.Tree[0].Text)
			}
		}

	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v\n", err)
	}

	return p

}

func makeCommand(line string) (Token, error) {

	line = strings.TrimSpace(line)

	if len(line) == 0 { // ignore empty lines
		return Token{}, nil
	} else if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") { // ignore comment lines
		return Token{}, nil
	}

	t := Token{
		Type: COMMAND,
		Text: line,
	}
	t.Tokenize()

	return t, nil
}

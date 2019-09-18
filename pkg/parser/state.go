package parser

import "fmt"

var State Stats

func init() {
	State = NewStats()
}

type Stats struct {
	Assertion Stat
	Must      Stat
	Should    Stat
}

func NewStats() Stats {
	return Stats{
		Assertion: NewStat(),
		Must:      NewStat(),
		Should:    NewStat(),
	}
}

type Stat struct {
	Failed    int
	Succeeded int
}

func NewStat() Stat {
	return Stat{
		Failed:    0,
		Succeeded: 0,
	}
}

func PrintStats() {
	//print stats here.
	fmt.Printf("%+v\n", State)
}

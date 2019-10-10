package parser

import (
	"time"
)

//A command to block the current thread
type SleepCommand struct {
	Millisecond int64
	When        Expression
}

func (sc *SleepCommand) Run(l *Lexer) interface{} {
	defer l.wg.Done()
	time.Sleep(time.Duration(sc.Millisecond) * time.Millisecond)
	return nil

}

func (sc *SleepCommand) SetWhen(expr Expression) {
	sc.When = expr
}

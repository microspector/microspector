package parser

import (
	"time"
)

//A command to block the current thread
type SleepCommand struct {
	Millisecond int64
}

func (sc *SleepCommand) Run(l *lex) interface{} {
	defer l.wg.Done()
	time.Sleep(time.Duration(sc.Millisecond) * time.Millisecond)
	return nil

}

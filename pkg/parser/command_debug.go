package parser

type DebugCommand struct {
}

func (hc *DebugCommand) Run() interface{} {
	return "we got a DEBUG Command here"
}

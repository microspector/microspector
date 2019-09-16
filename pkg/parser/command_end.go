package parser

type EndCommand struct {
}

func (hc *EndCommand) Run() interface{} {
	return "we got a SET Command here"
}
package parser

type AssertCommand struct {
}

func (hc *AssertCommand) Run() interface{} {
	return "we got an ASSERT Command here"
}

package parser

type ShouldCommand struct {
}

func (hc *ShouldCommand) Run() interface{} {
	return "we got a SHOULD Command here"
}

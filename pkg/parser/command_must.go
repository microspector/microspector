package parser

type MustCommand struct {
}

func (hc *MustCommand) Run() interface{} {
	return "we got a MUST Command here"
}

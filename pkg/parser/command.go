package parser

type State struct {
	Vars          map[string]interface{}
	SuccessShould int
	SuccessMust   int
	FailedShould  int
}

type Command interface {
	Run(State) error
}

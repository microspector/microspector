package command

type Command interface {
	run() interface{}
}

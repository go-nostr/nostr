package command

type Command interface {
	Name() string
	Init(args []string) error
	Run() error
}

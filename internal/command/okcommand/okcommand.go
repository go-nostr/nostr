package okcommand

import (
	"flag"
	"fmt"
)

func New() *OkCommand {
	cmd := &OkCommand{
		flagSet: flag.NewFlagSet("ok", flag.ContinueOnError),
	}
	return cmd
}

type OkCommand struct {
	flagSet *flag.FlagSet
}

func (c *OkCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *OkCommand) Name() string {
	return c.flagSet.Name()
}

func (c *OkCommand) Run() error {
	fmt.Print("oking!\n")
	return nil
}

package countcommand

import (
	"flag"
	"fmt"
)

func New() *CountCommand {
	cmd := &CountCommand{
		flagSet: flag.NewFlagSet("count", flag.ContinueOnError),
	}
	return cmd
}

type CountCommand struct {
	flagSet *flag.FlagSet
}

func (c *CountCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *CountCommand) Name() string {
	return c.flagSet.Name()
}

func (c *CountCommand) Run() error {
	fmt.Print("counting!\n")
	return nil
}

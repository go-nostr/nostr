package closecommand

import (
	"flag"
	"fmt"
)

func New() *CloseCommand {
	cmd := &CloseCommand{
		flagSet: flag.NewFlagSet("close", flag.ContinueOnError),
	}
	return cmd
}

type CloseCommand struct {
	flagSet *flag.FlagSet
}

func (c *CloseCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *CloseCommand) Name() string {
	return c.flagSet.Name()
}

func (c *CloseCommand) Run() error {
	fmt.Print("closing!\n")
	return nil
}

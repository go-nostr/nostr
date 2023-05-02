package authcommand

import (
	"flag"
	"fmt"
)

func New() *AuthCommand {
	cmd := &AuthCommand{
		flagSet: flag.NewFlagSet("auth", flag.ContinueOnError),
	}
	return cmd
}

type AuthCommand struct {
	flagSet *flag.FlagSet
}

func (c *AuthCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *AuthCommand) Name() string {
	return c.flagSet.Name()
}

func (c *AuthCommand) Run() error {
	fmt.Print("authorizing!\n")
	return nil
}

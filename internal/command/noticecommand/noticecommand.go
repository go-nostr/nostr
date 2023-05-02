package noticecommand

import (
	"flag"
	"fmt"
)

func New() *NoticeCommand {
	cmd := &NoticeCommand{
		flagSet: flag.NewFlagSet("notice", flag.ContinueOnError),
	}
	return cmd
}

type NoticeCommand struct {
	flagSet *flag.FlagSet
}

func (c *NoticeCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *NoticeCommand) Name() string {
	return c.flagSet.Name()
}

func (c *NoticeCommand) Run() error {
	fmt.Print("authorizing!\n")
	return nil
}

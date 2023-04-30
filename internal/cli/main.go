package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	CommandTypeHelp    = "help"
	CommandTypePublish = "publish"
)

var ()

func init() {
}

// nostr [subcommand] [options]
//
// where, subcommand∈{auth,close,count,event,req}

func main() {

	if len(os.Args[1:]) < 1 {
		flag.Usage = func() {
			fmt.Printf(
				"go-nostr: manage and communicate with Nostr protocol for decentralized social media." +
					"\n\n" +
					"Usage:" +
					"\n\n" +
					"\tnostr [command] [options]" +
					"\n\n" +
					"The commands are:" +
					"\n\n" +
					"\tauth\n" +
					"\tclose\n" +
					"\tcount\n" +
					"\tevent\n" +
					"\treq\n\n",
			)
		}
		flag.Usage()
		return
	}

	if err := os.Args[1:]; err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmds := []Command{
		&AuthCommand{},
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() != subcommand {
			fmt.Println(fmt.Errorf("unknown subcommand: %s", subcommand))
			return
		}
		cmd.Init(os.Args[2:])
		cmd.Run()
	}
}

type Command interface {
	Name() string
	Init(args []string) error
	Run() error
}

type BaseCommand struct {
	fs *flag.FlagSet
}

func (c *BaseCommand) Name() string {
	return c.fs.Name()
}

func (c *BaseCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *BaseCommand) Run() error {
	fmt.Printf("hello!")
	return nil
}

func NewAuthCommand() Command {
	ac := &AuthCommand{&BaseCommand{
		fs: flag.NewFlagSet("auth", flag.ContinueOnError),
	}}
	return ac
}

type AuthCommand struct {
	*BaseCommand
}

func NewCloseCommand() Command {
	ac := &CloseCommand{&BaseCommand{
		fs: flag.NewFlagSet("close", flag.ContinueOnError),
	}}
	return ac
}

type CloseCommand struct {
	*BaseCommand
}

type CountCommand struct {
	*BaseCommand
}

type EventCommand struct {
	*BaseCommand
}

type RequestCommand struct {
	*BaseCommand
}

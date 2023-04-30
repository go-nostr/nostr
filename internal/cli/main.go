package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr"
)

const (
	CommandTypeHelp    = "help"
	CommandTypePublish = "publish"
)

var cl = nostr.NewClient(nil)

func usageFunc() {
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

func init() {
	flag.Usage = usageFunc
}

func main() {
	if len(os.Args[1:]) < 1 {
		flag.Usage()
		os.Exit(1)
		return
	}

	cmds := map[string]Command{
		"auth":  NewAuthCommand(),
		"close": NewCloseCommand(),
		"count": NewCountCommand(),
		"event": NewEventCommand(),
		"req":   NewRequestCommand(cl),
	}

	subcommand := os.Args[1]

	if cmds[subcommand] == nil {
		flag.Usage()
		return
	}

	cmds[subcommand].Init(os.Args[2:])
	cmds[subcommand].Run()
}

type Command interface {
	Name() string
	Init(args []string) error
	Run() error
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

func (c *AuthCommand) Run() error {
	fmt.Print("authorizing-ng-ng!\n")
	return nil
}

func NewBaseCommand(name string) *BaseCommand {
	c := &BaseCommand{
		fs: flag.NewFlagSet(name, flag.ExitOnError),
	}
	c.fs.StringVar(&c.relay, "relay", "wss://nostr.wine", "Relay ...")
	return c
}

type BaseCommand struct {
	fs *flag.FlagSet

	relay string
}

func (c *BaseCommand) Name() string {
	return c.fs.Name()
}

func (c *BaseCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *BaseCommand) Run() error {
	fmt.Print("hello!\n")
	return nil
}

func NewCloseCommand() Command {
	cmd := &CloseCommand{&BaseCommand{
		fs: flag.NewFlagSet("close", flag.ContinueOnError),
	}}
	return cmd
}

type CloseCommand struct {
	*BaseCommand
}

func NewCountCommand() Command {
	cmd := &CountCommand{&BaseCommand{
		fs: flag.NewFlagSet("count", flag.ContinueOnError),
	}}
	return cmd
}

type CountCommand struct {
	*BaseCommand
}

func NewEventCommand() Command {
	cmd := &EventCommand{&BaseCommand{
		fs: flag.NewFlagSet("event", flag.ContinueOnError),
	}}
	return cmd
}

type EventCommand struct {
	*BaseCommand
}

func NewRequestCommand(cl *nostr.Client) Command {
	baseCommand := NewBaseCommand("req")
	cmd := &RequestCommand{
		BaseCommand: baseCommand,

		cl: cl,
	}

	cmd.fs.StringVar(&cmd.subscriptionID, "sid", "undefined", "Subscription ID used for ...")

	// TODO: add filterSlice to accept one or more filter, OR add one or more parameters to build filter...
	// cmd.fs.StringVar(&cmd.subscriptionID, "name", "World", "name of the person to be greeted")

	return cmd
}

type RequestCommand struct {
	*BaseCommand

	cl             *nostr.Client
	subscriptionID string
	// filter         *nostr.Filter
}

func (c *RequestCommand) Run() error {
	ctx := context.Background()
	content := make(chan []byte)
	mess := nostr.NewRequestMessage(c.subscriptionID, &nostr.Filter{})
	c.cl.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	c.cl.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	if err := c.cl.Subscribe(ctx, c.relay); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.cl.Publish(mess); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		select {
		case byt := <-content:
			fmt.Printf("%s\n\n", byt)
		case <-ctx.Done():
			os.Exit(0)
		}
	}
}

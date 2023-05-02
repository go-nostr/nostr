package command

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/go-nostr/nostr"
)

type Command interface {
	Name() string
	Init(args []string) error
	Run() error
}

func NewAuthCommand() *AuthCommand {
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
	cmd := &BaseCommand{
		fs: flag.NewFlagSet(name, flag.ExitOnError),
	}
	cmd.fs.StringVar(&cmd.u, "u", "wss://relay.damus.io", "Relay ...")
	cmd.fs.StringVar(&cmd.nsec, "nsec", "undefined", "Bech32 encoded private key ...")
	return cmd
}

type BaseCommand struct {
	fs *flag.FlagSet

	cl   *nostr.Client
	u    string
	nsec string
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

func NewCloseCommand() *CloseCommand {
	cmd := &CloseCommand{&BaseCommand{
		fs: flag.NewFlagSet("close", flag.ContinueOnError),
	}}
	return cmd
}

type CloseCommand struct {
	*BaseCommand
}

func NewCountCommand() *CountCommand {
	cmd := &CountCommand{&BaseCommand{
		fs: flag.NewFlagSet("count", flag.ContinueOnError),
	}}
	return cmd
}

type CountCommand struct {
	*BaseCommand
}

func NewNoticeCommand() *NoticeCommand {
	cmd := &NoticeCommand{&BaseCommand{
		fs: flag.NewFlagSet("notice", flag.ContinueOnError),
	}}
	return cmd
}

type NoticeCommand struct {
	*BaseCommand
}

func NewOkCommand() *OkCommand {
	cmd := &OkCommand{&BaseCommand{
		fs: flag.NewFlagSet("ok", flag.ContinueOnError),
	}}
	return cmd
}

type OkCommand struct {
	*BaseCommand
}

func NewEventCommand(cl *nostr.Client) *EventCommand {
	cmd := &EventCommand{
		BaseCommand: &BaseCommand{
			cl: cl,
			fs: flag.NewFlagSet("event", flag.ContinueOnError),
		},
	}
	cmd.fs.IntVar(&cmd.kind, "k", 0, "Event Kind ...")
	cmd.fs.StringVar(&cmd.content, "c", "", "Content ...")
	cmd.fs.StringVar(&cmd.u, "u", "undefined", "Subscription ID used for ...")
	cmd.fs.StringVar(&cmd.nsec, "nsec", "undefined", "Bech32 encoded private key ...")
	return cmd
}

type EventCommand struct {
	*BaseCommand

	content string
	kind    int
}

func (c *EventCommand) Run() error {
	ctx := context.Background()
	content := make(chan []byte)
	c.cl.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		fmt.Printf("%s:\t\t%s\n\n", mess.Values()[2].(map[string]any)["pubkey"], mess.Values()[2].(map[string]any)["content"])
	})
	c.cl.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	if err := c.cl.Subscribe(ctx, c.u); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, nsecBech, err := bech32.DecodeNoLimit(c.nsec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	privateKey, err := bech32.ConvertBits(nsecBech, 5, 8, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	privateKeyHex := hex.EncodeToString(privateKey[0:32])
	evnt := nostr.NewShortTextNoteEvent(c.content)
	evnt.Sign(privateKeyHex)
	mess := nostr.NewEventMessage("", evnt)
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

func NewRequestCommand(cl *nostr.Client) *RequestCommand {
	cmd := &RequestCommand{
		BaseCommand: &BaseCommand{
			cl: cl,
			fs: flag.NewFlagSet("event", flag.ContinueOnError),
		},
	}
	cmd.fs.StringVar(&cmd.subscriptionID, "sid", "undefined", "Subscription ID used for ...")
	cmd.fs.StringVar(&cmd.u, "u", "undefined", "Relay URL ...")
	return cmd
}

type RequestCommand struct {
	*BaseCommand

	subscriptionID string
	// TODO: add filter parameter parsing
	// filter         *nostr.Filter
}

func (c *RequestCommand) Run() error {
	ctx := context.Background()
	content := make(chan []byte)
	mess := nostr.NewRequestMessage(c.subscriptionID, &nostr.Filter{})
	c.cl.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		fmt.Printf("%s (%s):\t\t%s\n\n", mess.Values()[2].(map[string]any)["id"], mess.Values()[2].(map[string]any)["pubkey"], mess.Values()[2].(map[string]any)["content"])
	})
	c.cl.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	c.cl.HandleMessageFunc(nostr.MessageTypeOk, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	if err := c.cl.Subscribe(ctx, c.u); err != nil {
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

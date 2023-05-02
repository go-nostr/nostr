package requestcommand

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr"
)

func New(opt *Options) *RequestCommand {
	cmd := &RequestCommand{
		client:  opt.Client,
		flagSet: flag.NewFlagSet("req", flag.ExitOnError),
	}
	cmd.flagSet.StringVar(&cmd.subscriptionID, "sid", "undefined", "Subscription ID used for ...")
	cmd.flagSet.StringVar(&cmd.relay, "u", "undefined", "Relay URL ...")
	return cmd
}

type Options struct {
	Client *nostr.Client
}

type RequestCommand struct {
	client         *nostr.Client
	flagSet        *flag.FlagSet
	relay          string
	subscriptionID string
	// TODO: add filter parameter parsing
	// filter         *nostr.Filter
}

func (c *RequestCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *RequestCommand) Name() string {
	return c.flagSet.Name()
}

func (c *RequestCommand) Run() error {
	ctx := context.Background()
	content := make(chan []byte)
	mess := nostr.NewRequestMessage(c.subscriptionID, &nostr.Filter{})
	c.client.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		fmt.Printf("%s (%s):\t\t%s\n\n", mess.Values()[2].(map[string]any)["id"], mess.Values()[2].(map[string]any)["pubkey"], mess.Values()[2].(map[string]any)["content"])
	})
	c.client.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	c.client.HandleMessageFunc(nostr.MessageTypeOk, func(mess nostr.Message) {
		if data, err := mess.Marshal(); err == nil {
			content <- data
		}
	})
	if err := c.client.Subscribe(ctx, c.relay); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.client.Publish(mess); err != nil {
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

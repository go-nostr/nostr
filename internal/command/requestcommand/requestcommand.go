package requestcommand

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/requestmessage"
)

func New(opt *Options) *RequestCommand {
	cmd := &RequestCommand{
		client:  opt.Client,
		flagSet: flag.NewFlagSet("req", flag.ExitOnError),
	}
	cmd.flagSet.StringVar(&cmd.subscriptionID, "sid", "undefined", "Subscription ID used for ...")
	cmd.flagSet.StringVar(&cmd.relay, "u", "wss://relay.damus.io", "Relay URL ...")
	return cmd
}

type Options struct {
	Client *client.Client
}

type RequestCommand struct {
	client         *client.Client
	flagSet        *flag.FlagSet
	relay          string
	subscriptionID string
}

func (c *RequestCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *RequestCommand) Name() string {
	return c.flagSet.Name()
}

func (c *RequestCommand) Run() error {
	ctx := context.Background()
	c.client.HandleErrorFunc(func(err error) {
		fmt.Println(err.Error())
	})
	c.client.HandleMessageFunc(func(msg message.Message) {
		fmt.Println(msg.Values()...)
	})
	c.client.Connect(ctx, c.relay)
	msg := requestmessage.New(c.subscriptionID, &requestmessage.Filter{})
	c.client.SendMessage(ctx, msg)
	if err := c.client.Listen(ctx); err != nil {
		return err
	}
	return nil
}

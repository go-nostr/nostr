package requestcommand

import (
	"context"
	"flag"
	"fmt"
	"os"

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
	content := make(chan []byte)
	mess := requestmessage.New(c.subscriptionID, &requestmessage.Filter{})
	c.client.HandleMessageFunc(func(mess *message.Message) {
		if mess.Values()[0].(string) == "EVENT" {
			fmt.Printf("%s:\n\n%s\n\n", mess.Values()[2].(map[string]any)["pubkey"], mess.Values()[2].(map[string]any)["content"])
		}
	})
	if err := c.client.Subscribe(ctx, c.relay); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.client.Publish(ctx, mess); err != nil {
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

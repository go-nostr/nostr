package eventcommand

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/event/shorttextnote"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/eventmessage"
)

func New(opt *Options) *EventCommand {
	cmd := &EventCommand{
		Options: opt,
		fs:      flag.NewFlagSet("event", flag.ContinueOnError),
	}
	cmd.fs.IntVar(&cmd.Kind, "k", 0, "Event Kind ...")
	cmd.fs.StringVar(&cmd.Content, "c", "", "Content ...")
	cmd.fs.StringVar(&cmd.Relay, "u", "undefined", "Subscription ID used for ...")
	cmd.fs.StringVar(&cmd.Nsec, "nsec", "undefined", "Bech32 encoded private key ...")
	return cmd
}

type Options struct {
	Client  *client.Client
	Content string
	Kind    int
	Nsec    string
	Relay   string
}

type EventCommand struct {
	*Options

	fs *flag.FlagSet
}

func (c *EventCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *EventCommand) Name() string {
	return c.fs.Name()
}

func (c *EventCommand) Run() error {
	ctx := context.Background()
	errChan := make(chan error)
	msgChan := make(chan message.Message)
	c.Client.HandleErrorFunc(func(err error) {
		errChan <- err
	})
	c.Client.HandleMessageFunc(func(msg message.Message) {
		msgChan <- msg
	})
	c.Client.Connect(ctx, c.Relay)
	_, nsecData, err := bech32.DecodeNoLimit(c.Nsec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prvKey, err := bech32.ConvertBits(nsecData, 5, 8, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prvKeyHex := hex.EncodeToString(prvKey[0:32])
	evt := shorttextnote.New(c.Content)
	evt.Sign(prvKeyHex)
	outMsg := eventmessage.New("", evt)
	c.Client.SendMessage(ctx, outMsg)
	inMsg := <-msgChan
	inMsgByt, err := inMsg.Marshal()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", inMsgByt)
	return nil
}

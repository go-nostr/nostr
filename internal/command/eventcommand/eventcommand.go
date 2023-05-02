package eventcommand

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/go-nostr/nostr"
)

func New(opt *Options) *EventCommand {
	cmd := &EventCommand{
		client:  opt.Client,
		flagSet: flag.NewFlagSet("event", flag.ContinueOnError),
	}
	cmd.flagSet.IntVar(&cmd.kind, "k", 0, "Event Kind ...")
	cmd.flagSet.StringVar(&cmd.content, "c", "", "Content ...")
	cmd.flagSet.StringVar(&cmd.relay, "u", "undefined", "Subscription ID used for ...")
	cmd.flagSet.StringVar(&cmd.nsec, "nsec", "undefined", "Bech32 encoded private key ...")
	return cmd
}

type Options struct {
	Client *nostr.Client
}

type EventCommand struct {
	content string
	client  *nostr.Client
	flagSet *flag.FlagSet
	kind    int
	nsec    string
	relay   string
}

func (c *EventCommand) Init(args []string) error {
	return c.flagSet.Parse(args)
}

func (c *EventCommand) Name() string {
	return c.flagSet.Name()
}

func (c *EventCommand) Run() error {
	ctx := context.Background()
	messChan := make(chan nostr.Message)
	c.client.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		messChan <- mess
	})
	c.client.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		messChan <- mess
	})
	c.client.HandleMessageFunc(nostr.MessageTypeOk, func(mess nostr.Message) {
		messChan <- mess
	})
	if err := c.client.Subscribe(ctx, c.relay); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, nsecBech, err := bech32.DecodeNoLimit(c.nsec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prvKey, err := bech32.ConvertBits(nsecBech, 5, 8, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prvKeyHex := hex.EncodeToString(prvKey[0:32])
	evnt := nostr.NewShortTextNoteEvent(c.content)
	evnt.Sign(prvKeyHex)
	outMess := nostr.NewEventMessage("", evnt)
	if err := c.client.Publish(outMess); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inMess := <-messChan
	inMessByt, err := inMess.Marshal()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", inMessByt)
	return nil
}

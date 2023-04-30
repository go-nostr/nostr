package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/go-nostr/nostr"
)

const (
	CommandTypeHelp    = "help"
	CommandTypePublish = "publish"
)

type relaySlice []string

func (rs *relaySlice) String() string {
	return strings.Join(*rs, ",")
}

func (rs *relaySlice) Set(u string) error {
	*rs = append(*rs, u)
	return nil
}

var (
	data   string     = "[]"
	relays relaySlice = make(relaySlice, 0)
)

func init() {
	flag.StringVar(&data, "d", data, "")
	flag.Var(&relays, "r", "")
}

func main() {
	flag.Parse()

	ch := make(chan nostr.Message)
	cl := nostr.NewClient(nil)

	cl.HandleMessageFunc(nostr.MessageTypeEOSE, func(mess nostr.Message) {
		ch <- mess
	})
	cl.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		ch <- mess
	})
	cl.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		ch <- mess
	})
	cl.HandleMessageFunc(nostr.MessageTypeRequest, func(mess nostr.Message) {
		ch <- mess
	})

	ctx := context.Background()

	for _, u := range relays {
		fmt.Printf("Connecting... (%s)", u)
		if err := cl.Subscribe(ctx, u); err != nil {
			panic(err)
		}
	}

	fmt.Printf(" Connected!\n")

	mess := nostr.NewRequestMessage("asdf-1234", &nostr.Filter{})
	byt, _ := mess.Marshal()
	fmt.Printf("%s", byt)
	go cl.Publish(mess)

	for {
		select {
		case mess := <-ch:
			if string(mess.Type()) == nostr.MessageTypeEvent {
				eventMess := &nostr.EventMessage{RawMessage: mess.(*nostr.RawMessage)}
				fmt.Printf("\nFrom: %s\n\n%s\n\n\n", eventMess.Event().PublicKey(), eventMess.Event().Content())
			}
		case <-ctx.Done():
			return
		}
	}
}

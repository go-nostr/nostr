package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr"
)

// Initialize a new Nostr client with default settings, and set default relay and subscription ID values
var cl = nostr.NewClient(nil)
var relay string = "wss://relay.damus.io"
var sid string = "asdf-1234"

// Parse command line arguments for custom relay and subscription ID values
func init() {
	flag.StringVar(&relay, "relay", relay, "Relay to connect to")
	flag.StringVar(&sid, "sid", sid, "Subscription ID to use for connection")
}

func main() {
	// Parse command line flags, allowing users to customize relay and subscription ID values
	flag.Parse()
	// Create a context for managing the lifecycle of the Nostr client
	ctx := context.Background()
	// Subscribe to the specified relay
	cl.Subscribe(ctx, relay)
	// Publish a request message to the relay to establish connection
	cl.Publish(nostr.NewRequestMessage(sid, &nostr.Filter{}))
	// Handle the End Of Subscription Epoch (EOSE) message, which indicates the end of a subscription epoch
	cl.HandleMessageFunc(nostr.MessageTypeEOSE, func(mess nostr.Message) {
		vals := mess.Values()
		sid, ok := vals[1].(string)
		if !ok {
			return
		}
		fmt.Printf("EOSE:\t%s\n\n", sid)
	})
	// Handle Event messages, which are used for sending and receiving events in the Nostr network
	cl.HandleMessageFunc(nostr.MessageTypeEvent, func(mess nostr.Message) {
		vals := mess.Values()
		event, ok := vals[2].(map[string]interface{})
		if !ok {
			return
		}
		if event["kind"].(float64) == 1 {
			fmt.Printf("npub:%s\n\n%s\n\n\n", event["pubkey"], event["content"])
		}
	})
	// Handle Notice messages, which are used for sending and receiving notices in the Nostr network
	cl.HandleMessageFunc(nostr.MessageTypeNotice, func(mess nostr.Message) {
		vals := mess.Values()
		notice, ok := vals[1].(map[string]string)
		if !ok {
			return
		}
		fmt.Printf("NOTICE:\t%s\n\n", notice)
	})
	// Continuously listen for incoming messages until the context is cancelled
	for {
		select {
		case <-ctx.Done():
			os.Exit(0)
		}
	}
}

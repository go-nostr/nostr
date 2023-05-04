package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/requestmessage"
)

// Initialize a new Nostr client with default settings, and set default relay and subscription ID values
var cl = client.New(nil)
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
	cl.Publish(ctx, requestmessage.New(sid, &requestmessage.Filter{}))
	// Handle the End Of Subscription Epoch (EOSE) message, which indicates the end of a subscription epoch
	cl.HandleMessageFunc(func(mess *message.Message) {
		vals := mess.Values()
		sid, ok := vals[1].(string)
		if !ok {
			return
		}
		fmt.Printf("EOSE:\t%s\n\n", sid)
	})
	// Continuously listen for incoming messages until the context is cancelled
	if err := cl.Listen(ctx); err != nil {
		fmt.Println("%w", err)
		os.Exit(1)
	}
}

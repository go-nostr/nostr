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
var u string = "wss://relay.damus.io"
var sid string = "asdf-1234"

// Parse command line arguments for custom relay and subscription ID values
func init() {
	flag.StringVar(&u, "relay", u, "Relay to connect to")
	flag.StringVar(&sid, "sid", sid, "Subscription ID to use for connection")
}

func main() {
	// Parse command line flags, allowing users to customize relay and subscription ID values
	flag.Parse()
	// Create a context for managing the lifecycle of the Nostr client
	ctx := context.Background()
	cl := client.New(&client.Options{
		ReadLimit: 2e6,
	})
	// Connect to the specified relay
	cl.Connect(ctx, u)
	// Send a request message to the relay to establish connection
	cl.HandleErrorFunc(func(err error) {
		fmt.Println(err.Error())
	})
	cl.HandleMessageFunc(func(msg message.Message) {
		fmt.Println(msg)
	})
	cl.SendMessage(ctx, requestmessage.New(sid, &requestmessage.Filter{}))
	// Handle the End Of Subscription Epoch (EOSE) message, which indicates the end of a subscription epoch
	// Continuously listen for incoming messages until the context is cancelled
	if err := cl.Listen(ctx); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

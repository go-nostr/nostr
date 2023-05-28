package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/requestmessage"
	"github.com/go-nostr/nostr/subscriptionid"
)

var (
	// Declare and initialize a string variable 'relayURL' which will be used
	// as the default URL for the relay server.
	relayURL string = "wss://relay.damus.io"

	// Declare and initialize a string variable 'subscriptionID' which will be
	// used as the default Subscription ID.
	subscriptionID string = subscriptionid.New()
)

func init() {
	// Register a command-line option "-u" with a default value of 'relayURL'
	// and a description. This will allow the user to provide a custom relay URL.
	flag.StringVar(&relayURL, "u", relayURL, "Relay to connect to")

	// Register a command-line option "-sid" with a default value of
	// 'subscriptionID' and a description. This will allow the user to provide
	// a custom subscription ID.
	flag.StringVar(&subscriptionID, "sid", subscriptionID, "Subscription ID to use for connection")
}

func main() {
	// Parse the command-line options provided by the user. The parsed options
	// will update the values of 'relayURL' and 'subscriptionID' if they were
	// provided.
	flag.Parse()

	// Create a new context. This context will be used for various operations
	// like connecting to the relay, sending a message, and listening for
	// incoming messages.
	ctx := context.Background()

	// Create a new Nostr client with a specified read limit.
	cl := client.New(&client.Options{
		ReadLimit: 2e6,
	})

	// Connect the Nostr client to the relay server using the provided URL.
	cl.Connect(ctx, relayURL)

	// Register an error handling function that will print out any errors that
	// occur while the client is running.
	cl.HandleErrorFunc(func(err error) {
		fmt.Println(err.Error())
	})

	// Register a message handling function that will print out any messages
	// that are received by the client.
	cl.HandleMessageFunc(func(msg message.Message) {
		fmt.Println(msg)
	})

	// Send a subscription message to the relay server using the provided
	// Subscription ID. The subscription message has an empty filter which
	// means it will receive all messages.
	cl.SendMessage(ctx, requestmessage.New(subscriptionID, &requestmessage.Filter{}))

	// Start listening for incoming messages. If any error occurs while
	// listening, print the error and exit the program with a non-zero status
	// code.
	if err := cl.Listen(ctx); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

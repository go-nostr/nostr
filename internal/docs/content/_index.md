---
title: go-nostr/nostr
---

go-nostr/nostr is a software component of the Nostr network, a decentralized and extensible social protocol. This piece of software serves as a bridge between Nostr-compatible applications, offering both relay and client functionalities to ensure smooth communication within the network.

As an integral part of the Nostr network, go-nostr/nostr aids in the creation, management, and communication of events between clients and relays. It provides core functionalities and supports selected NIPs (Nostr Improvement Proposals), simplifying the development process of Nostr-compatible applications and contributing to the overall growth of the network.

## Getting Started

The go-nostr project is structured into three main sections:

1. **Go internal executable and dependencies**: These are the core elements that run the go-nostr software.
2. **Go external package**: This includes essential components required for building custom clients and relays.
3. **Angular web application**: This facilitates the administration of the client and interaction with third-party clients and relays.

Here's a simple go-nostr usage example using the Go package to create a simple Nostr client:

```go
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
 relayURL string = "wss://relay.damus.io"
 subscriptionID string = subscriptionid.New()
)

func init() {
 flag.StringVar(&relayURL, "u", relayURL, "Relay to connect to")
 flag.StringVar(&subscriptionID, "sid", subscriptionID, "Subscription ID to use for connection")
}

func main() {
 flag.Parse()

 ctx := context.Background()

 cl := client.New(&client.Options{
  ReadLimit: 2e6,
 })

 cl.Connect(ctx, relayURL)

 cl.HandleErrorFunc(func(err error) {
  fmt.Println(err.Error())
 })

 cl.HandleMessageFunc(func(msg message.Message) {
  fmt.Println(msg)
 })

 cl.SendMessage(ctx, requestmessage.New(subscriptionID, &requestmessage.Filter{}))

 if err := cl.Listen(ctx); err != nil {
  fmt.Print(err.Error())
  os.Exit(1)
 }
}

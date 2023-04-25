//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/internal/web"
	"github.com/google/wire"
)

const (
	defaultHostname   = "0.0.0.0"
	defaultClientPort = 4200
	defaultRelayPort  = 3001
)

func provideClientAddr() string {
	return fmt.Sprintf("%v:%v", defaultHostname, defaultClientPort)
}

func provideRelayAddr() string {
	return fmt.Sprintf("%v:%v", defaultHostname, defaultRelayPort)
}

func provideClientHandler() http.Handler {
	dist, err := fs.Sub(web.FS, "dist")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	return http.FileServer(http.FS(dist))
}

func provideRelayHandler() http.Handler {
	return nostr.NewRelay()
}

func buildClientServer() *http.Server {
	wire.Build(wire.NewSet(
		provideClientAddr,
		provideClientHandler,
		wire.Struct(new(http.Server), "Addr", "Handler"),
	))
	return &http.Server{}
}

func buildRelayServer() *http.Server {
	wire.Build(wire.NewSet(
		provideRelayAddr,
		provideRelayHandler,
		wire.Struct(new(http.Server), "Addr", "Handler"),
	))
	return &http.Server{}
}

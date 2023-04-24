//go:build wireinject
// +build wireinject

package main

import (
	"io/fs"
	"net/http"

	"github.com/go-nostr/nostr/internal/web"
	"github.com/google/wire"
)

func provideFileServerHandler() http.Handler {
	dist, _ := fs.Sub(web.FS, "dist")
	return http.FileServer(http.FS(dist))
}

func buildHTTPServer() *http.Server {
	wire.Build(wire.NewSet(
		provideFileServerHandler,
		wire.Struct(new(http.Server), "Handler"),
	))
	return &http.Server{}
}

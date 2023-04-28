package main

import (
	"context"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	// NOTE: Create a new context for managing the lifetime of the main
	// function
	ctx := context.Background()
	// NOTE: Initialize a struct containing all the service servers
	sc := struct {
		cl *http.Server
		rl *http.Server
	}{
		cl: buildClientServer(),
		rl: buildRelayServer(),
	}
	// NOTE: Create an error group to manage the concurrent execution of
	// service servers
	g, _ := errgroup.WithContext(ctx)
	// NOTE: Add service server Serve functions to the error group
	g.Go(sc.cl.ListenAndServe)
	g.Go(sc.rl.ListenAndServe)
	// NOTE: Wait for all service servers to complete or return an error
	if err := g.Wait(); err != nil {
		panic(err)
	}
}

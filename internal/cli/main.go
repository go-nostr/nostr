package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr/internal/command"
)

func usageFunc() {
	fmt.Printf(
		"go-nostr: manage and communicate with Nostr protocol for decentralized social media." +
			"\n\n" +
			"Usage:" +
			"\n\n" +
			"\tnostr [command] [options]" +
			"\n\n" +
			"The commands are:" +
			"\n\n" +
			"\tauth\n" +
			"\tclose\n" +
			"\tcount\n" +
			"\tevent\n" +
			"\treq\n\n",
	)
}

func init() {
	flag.Usage = usageFunc
}

func main() {
	if len(os.Args[1:]) < 1 {
		flag.Usage()
		os.Exit(1)
		return
	}
	cmds := map[string]command.Command{
		"auth":  buildAuthCommand(),
		"close": buildCloseCommand(),
		"count": buildCountCommand(),
		"event": buildEventCommand(),
		"req":   buildRequestCommand(),
	}
	name := os.Args[1]
	if cmds[name] == nil {
		flag.Usage()
		return
	}
	cmds[name].Init(os.Args[2:])
	cmds[name].Run()
}

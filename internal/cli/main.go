package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-nostr/nostr/internal/command"
)

func usageFunc() {
	fmt.Printf("go-nostr is a tool for communicating over the Nostr protocol.\n" +
		"\n" +
		"Usage:\n" +
		"\n" +
		"\tnostr [command] [options]\n" +
		"\n" +
		"The commands are:\n" +
		"\n" +
		"\tauth\tused to send authentication events\n" +
		"\tclose\tused to stop previous subscriptions\n" +
		"\tcount\tused to request event counts\n" +
		"\tevent\tused to publish events\n" +
		"\tnotice\tused to send human-readable messages to clients\n" +
		"\tok\tused to notify clients if an EVENT was successful\n" +
		"\treq\tused to request events and subscribe to new updates\n" +
		"\n" +
		"Use \"nostr help <command>\" for more information about a command.\n" +
		"\n",
	)
}

func init() {
	flag.Usage = usageFunc
}

func main() {
	if len(os.Args[1:]) < 1 {
		flag.Usage()
		os.Exit(0)
	}
	cmds := map[string]command.Command{
		"auth":   buildAuthCommand(),
		"close":  buildCloseCommand(),
		"count":  buildCountCommand(),
		"notice": buildNoticeCommand(),
		"ok":     buildOkCommand(),
		"event":  buildEventCommand(),
		"req":    buildRequestCommand(),
	}
	name := os.Args[1]
	if cmds[name] == nil {
		flag.Usage()
		os.Exit(0)
	}
	cmds[name].Init(os.Args[2:])
	cmds[name].Run()
}

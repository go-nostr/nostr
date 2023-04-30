//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/internal/command"
	"github.com/google/wire"
)

func buildClient() *nostr.Client {
	return nostr.NewClient(nil)
}

func buildAuthCommand() *command.AuthCommand {
	wire.Build(wire.NewSet(
		command.NewAuthCommand,
	))
	return &command.AuthCommand{}
}

func buildCountCommand() *command.CountCommand {
	wire.Build(
		command.NewCountCommand,
	)
	return &command.CountCommand{}
}

func buildCloseCommand() *command.CloseCommand {
	wire.Build(
		command.NewCloseCommand,
	)
	return &command.CloseCommand{}

}

func buildEventCommand() *command.EventCommand {
	wire.Build(
		command.NewEventCommand,
	)
	return &command.EventCommand{}
}

func buildNoticeCommand() *command.NoticeCommand {
	wire.Build(
		command.NewNoticeCommand,
	)
	return &command.NoticeCommand{}

}

func buildRequestCommand() *command.RequestCommand {
	wire.Build(
		buildClient,
		command.NewRequestCommand,
	)
	return &command.RequestCommand{}

}

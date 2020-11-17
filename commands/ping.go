package commands

import "github.com/Necroforger/dgrouter/disgordrouter"

func pingCommand(ctx *disgordrouter.Context) {
	ctx.Reply("Pong!")
}

func init() {
	NewCommand("ping", "responds with pong", "General", pingCommand)
}

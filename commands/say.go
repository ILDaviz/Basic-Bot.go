package commands

import (
	"github.com/Necroforger/dgrouter/disgordrouter"
	"strings"
)

func sayCommand(ctx *disgordrouter.Context) {
	if ctx.Args.Get(1) != "" {
		/*
			<prefix>say -> Args[0]
			<prefix>say text -> Args[1]
		*/
		ctx.Reply(strings.Join(ctx.Args[1:], " "))

	} else {
		ctx.Reply("There are no arguments to send")
	}
}

func init() {
	NewCommand("say", "Repeat the content you put", "General", sayCommand)
}

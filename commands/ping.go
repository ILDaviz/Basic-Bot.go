package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/Necroforger/dgrouter/disgordrouter"
)

func pingCommand(ctx *disgordrouter.Context) {
	start := time.Now()
	msg, err := ctx.Reply("Pong!")
	if err != nil {
		log.Fatal(err)
	} else {
		took := time.Since(start)
		_, _ = ctx.Ses.Channel(ctx.Msg.ChannelID).Message(msg.ID).SetContent(fmt.Sprintf("Pong! `%s`", took))
	}
}

func init() {
	NewCommand("ping", "Respond with pong", "General", pingCommand)
}

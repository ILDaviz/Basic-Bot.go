package commands

import (
	"context"
	"fmt"
	"github.com/Necroforger/dgrouter/disgordrouter"
	"log"
	"time"
)

func pingCommand(ctx *disgordrouter.Context) {
	start := time.Now()
	msg, err := ctx.Reply("Pong!")
	if err != nil {
		log.Fatal(err)
	} else {
		took := time.Since(start)
		ctx.Ses.SetMsgContent(context.Background(), ctx.Msg.ChannelID, msg.ID, fmt.Sprintf("Pong! `%s`", took))
	}
}

func init() {
	NewCommand("ping", "Respond with pong", "General", pingCommand)
}

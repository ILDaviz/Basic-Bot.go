package commands

import (
	"fmt"
	"strings"

	"github.com/Necroforger/dgrouter/disgordrouter"

	"Bot/utils"
)

func kickCommand(ctx *disgordrouter.Context) {
	PermsBot := util.HasPermission(ctx.Ses, BotUser, ctx.Msg.GuildID, util.KICK_MEMBERS) //Check the "Kick Members" Permission for the bot

	PermsUser := util.HasPermission(ctx.Ses, ctx.Msg.Author, ctx.Msg.GuildID, util.KICK_MEMBERS) //Check the "Kick members" permission for the ctx.Msg author

	if !PermsUser {
		ctx.Reply("Sorry, but you don't have permission to kick members")
		return
	}

	if !PermsBot {
		ctx.Reply("I don't have permission to kick members")
		return
	}

	if len(ctx.Msg.Mentions) < 1 {
		ctx.Reply("You must mention a member to kick")
		return
	}

	if ctx.Args.Get(2) != "" {
		err := ctx.Ses.Guild(ctx.Msg.GuildID).Member(ctx.Msg.Mentions[0].ID).Kick(strings.Join(ctx.Args[2:], " "))

		if err != nil {
			ctx.Reply(fmt.Sprintf("an unexpected error occurred: `%s`", err))
			return
		}

		ctx.Reply(fmt.Sprintf("`%s`, Successfully kicked!\n> **Reason**: %s", ctx.Msg.Mentions[0].Tag(), strings.Join(ctx.Args[2:], " ")))

	} else {
		ctx.Reply("You must add a reason for the kick")
	}
}

func init() {
	NewCommand("kick", "Kick a member from the server", "Moderation", kickCommand)
}

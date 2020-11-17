package commands

import (
	"Bot/utils"
	"context"
	"fmt"
	"github.com/Necroforger/dgrouter/disgordrouter"
	"github.com/andersfylling/disgord"
	"strings"
)

func banCommand(ctx *disgordrouter.Context) {
	PermsBot := util.HasPermission(ctx.Ses, BotUser, ctx.Msg.GuildID, util.BAN_MEMBERS) //Check the "Ban Members" Permission for the bot

	PermsUser := util.HasPermission(ctx.Ses, ctx.Msg.Author, ctx.Msg.GuildID, util.BAN_MEMBERS) //Check the "Ban members" permission for the ctx.Msg author

	if !PermsUser {
		ctx.Reply("Sorry, but you don't have permission to ban members")
		return //If you do not have permission, you are notified and execution ends
	}

	if !PermsBot {
		ctx.Reply("I don't have permission to ban members")
		return //If the bot does not have permission, it is notified and execution is terminated
	}

	if len(ctx.Msg.Mentions) < 1 {
		//If there was no mention to any user
		ctx.Reply("You must mention a member to ban")

	} else {

		if ctx.Args.Get(2) != "" {

			err := ctx.Ses.BanMember(context.Background(), ctx.Msg.GuildID, ctx.Msg.Mentions[0].ID, &disgord.BanMemberParams{DeleteMessageDays: 7, Reason: strings.Join(ctx.Args[2:], " ")}, disgord.IgnoreCache)

			if err != nil {
				ctx.Reply(fmt.Sprintf("an unexpected error occurred: `%s`", err))
				return //If an error occurs during the ban
			}

			ctx.Reply(fmt.Sprintf("`%s`, Successfully banned!\n> **Reason:** %s", ctx.Msg.Mentions[0].Tag(), strings.Join(ctx.Args[2:], " ")))
			//Ban executed successfully
		} else {
			//If there was no reason for the ban
			ctx.Reply("You must add a reason for the ban")
		}
	}
}

func init() {
	NewCommand("ban", "Ban a member from the server", "Moderation", banCommand)
}

package commands

import (
	"Bot/utils"
	"context"
	"github.com/Necroforger/dgrouter/disgordrouter"
)

func embedCommand(ctx *disgordrouter.Context) {
	avatar, _ := ctx.Msg.Author.AvatarURL(2048, true)
	avatarBot, _ := BotUser.AvatarURL(2048, false)

	embed := util.NewEmbed().
		SetTitle("The embed title").
		SetDescription("The embed description").
		SetUrl("https://github.com/Danny2105").
		SetColor(0xe9e9e9).
		SetTimestamp().
		SetFooter("Footer text", avatarBot).
		SetAuthor(ctx.Msg.Author.Tag(), avatar).
		AddField("Field name", "Field Value", true).
		AddField("Field name 2", "Field value 2", true).
		AddField("Field name 3", "Field value 3", true).
		AddField("Field name 4", "Field value 4", true).
		AddField("Field name no inline", "Field value no inline", false)

	_, _ = ctx.Ses.CreateMessage(context.Background(), ctx.Msg.ChannelID, embed.ToMessage())
}

func init() {
	NewCommand("embed", "Show a embed!", "General", embedCommand)
}

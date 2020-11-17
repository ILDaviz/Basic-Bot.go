package commands

import (
	"Bot/utils"
	"context"
	"fmt"
	"github.com/Necroforger/dgrouter/disgordrouter"
)

func avatarCommand(ctx *disgordrouter.Context) {
	mention := ctx.Msg.Mentions
	avatarEmbed := util.NewEmbed().SetTimestamp().SetColor(0xe9e9e9)
	avatarAuthor, _ := ctx.Msg.Author.AvatarURL(2048, true)

	if len(mention) > 0 {
		/* If there was more than 1 mention, we only get the first mention */
		mAvatar, _ := mention[0].AvatarURL(2048, true)

		avatarEmbed.
			SetTitle(fmt.Sprintf("Avatar of %s", mention[0].Tag())).
			SetFooter(fmt.Sprintf("Request by: %s", ctx.Msg.Author.Username), avatarAuthor).
			SetImage(mAvatar).
			SetAuthor(mention[0].Username, mAvatar)

		_, _ = ctx.Ses.CreateMessage(context.Background(), ctx.Msg.ChannelID, avatarEmbed.ToMessage())

	} else {
		/* If there was no mention send that of the author */
		avatarEmbed.
			SetImage(avatarAuthor)

		_, _ = ctx.Ses.CreateMessage(context.Background(), ctx.Msg.ChannelID, avatarEmbed.ToMessage())
	}

}

func init() {
	NewCommand("avatar", "Show someone else's avatar or yours!", "General", avatarCommand)
}

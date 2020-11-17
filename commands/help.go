package commands

import (
	"Bot/utils"
	"context"
	"fmt"
	"github.com/Necroforger/dgrouter/disgordrouter"
	"strings"
)

func helpCommand(ctx *disgordrouter.Context) {
	avatarAuthor, _ := ctx.Msg.Author.AvatarURL(2048, true)
	embed := util.NewEmbed().
		SetTimestamp().
		SetColor(0xe9e9e9).
		SetFooter(fmt.Sprintf("Request by: %s", ctx.Msg.Author.Username), avatarAuthor)

	if ctx.Args.Get(1) == "" {

		for category, commands := range getCategories() {
			embed.AddField(category, strings.Join(commands, ", "), false)
		}
	} else {
		if cmd, exist := Commands[ctx.Args.Get(1)]; exist {
			embed.
				SetTitle(fmt.Sprintf("Detailed help for %s", cmd.Name)).
				AddField("Name", cmd.Name, true).
				AddField("Category", cmd.Category, true).
				AddField("Description", cmd.Description, false)
		} else {
			ctx.Reply(fmt.Sprintf("No Command found with name `%s`", ctx.Args.Get(1)))
			return
		}
	}

	_, _ = ctx.Ses.CreateMessage(context.Background(), ctx.Msg.ChannelID, embed.ToMessage())
}

func getCategories() map[string][]string {
	var categories = make(map[string][]string)

	for _, cmd := range Commands {
		if cat, exist := categories[cmd.Category]; exist {
			cat = append(cat, fmt.Sprintf("`%s`", cmd.Name))
			categories[cmd.Category] = cat
		} else {
			categories[cmd.Category] = []string{fmt.Sprintf("`%s`", cmd.Name)}
		}
	}

	return categories
}

func init() {
	NewCommand("help", "Show this help message", "General", helpCommand)
}

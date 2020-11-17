package commands

import (
	"Bot/database"
	"Bot/utils"
	"fmt"
	"github.com/Necroforger/dgrouter/disgordrouter"
)

func setPrefixCommand(ctx *disgordrouter.Context) {
	if !util.HasPermission(ctx.Ses, ctx.Msg.Author, ctx.Msg.GuildID, util.MANAGE_GUILD) {
		ctx.Reply("You do not have permission to change the server prefix")

	} else {
		if ctx.Args.Get(1) == "" {
			ctx.Reply("Enter the new prefix")

		} else {
			newPrefix := ctx.Args.Get(1)

			if len(newPrefix) > 5 {
				ctx.Reply("The maximum prefix length is 5 and the minimum is 1")
				return
			}

			err := database.Run("UPDATE settings SET Prefix = ? WHERE GuildId = ?", newPrefix, ctx.Msg.GuildID)

			if err != nil {
				ctx.Reply(fmt.Sprintf("An error occurred while changing the prefix\nError: `%s`", err))
				return
			}

			ctx.Reply(fmt.Sprintf("Prefix updated to `%s`", newPrefix))
		}
	}
}

func init() {
	NewCommand("setprefix", "Set the custom prefix", "Administration", setPrefixCommand)
}

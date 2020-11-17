package commands

import (
	"Bot/utils"
	"context"
	"fmt"
	"github.com/Necroforger/dgrouter/disgordrouter"
	"github.com/andersfylling/disgord"
	"strconv"
)

func clearCommand(ctx *disgordrouter.Context) {
	if !util.HasPermission(ctx.Ses, ctx.Msg.Author, ctx.Msg.GuildID, util.MANAGE_MESSAGES) {
		ctx.Reply("You do not have the `Manage messages` permission")

	} else if !util.HasPermission(ctx.Ses, BotUser, ctx.Msg.GuildID, util.MANAGE_MESSAGES) {
		ctx.Reply("I don't have the `Manage messages` permission")

	} else if ctx.Args.Get(1) == "" {
		ctx.Reply("You must enter the number of messages to remove")

	} else {

		num, err := strconv.Atoi(ctx.Args.Get(1))
		//ctx.Args.Get(1) <- string
		//num -> int

		if err != nil {
			//If error is different from nil, it means that the amount entered is not a number
			ctx.Reply("You must enter only numbers")
			return
		}

		if num > 100 || num < 1 {
			ctx.Reply("The range to delete messages is **1 - 100**")

		} else {

			//If number is greater than 0 and less than or equal to 100 we continue
			Params := &disgord.GetMessagesParams{
				Limit: uint(num),
			}

			MessagesAmount, err := ctx.Ses.GetMessages(context.Background(), ctx.Msg.ChannelID, Params)
			//We get the number of Messages
			//MessageAmount []*disgord.Message

			if err != nil {
				ctx.Reply(fmt.Sprintf("An error occurred to verify the messages to delete\nError: %s", err))
				return
			}
			//Type disgord.DeleteMessagesParams
			deleteParams := &disgord.DeleteMessagesParams{}

			for i := 0; i < len(MessagesAmount); i++ {
				//This for is in charge of adding the messages that are going to be removed from the channel
				deleteParams.AddMessage(MessagesAmount[i])
				/*(Params *disgord.DeleteMessagesParams) AddMessage(msg *disgord.Message)*/
			}

			if err = ctx.Ses.DeleteMessages(context.Background(), ctx.Msg.ChannelID, deleteParams, disgord.IgnoreCache); err != nil {
				ctx.Reply(fmt.Sprintf("They cannot be removed `%d` messages\nError: `%s`", num, err))
				return
				//If error is different from nil it means that there was an error in deleting the indicated number of messages (usually because messages of more than 2 weeks cannot be deleted)
			}

			ctx.Reply(fmt.Sprintf("I just deleted `%d` message(s)", num))
			//If everything went well, I delete the messages and it is reported that the number of messages indicated was successfully deleted.
		}
	}
}

func init() {
	NewCommand("clear", "Clear messages", "Moderation", clearCommand)
}

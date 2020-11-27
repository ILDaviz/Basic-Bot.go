package main //package main (main file)

import (
	"log" // package log

	"github.com/Necroforger/dgrouter/disgordrouter" // command router framework
	"github.com/andersfylling/disgord"              // lib disgord

	"Bot/commands"
	"Bot/config"
)

func main() {
	client := disgord.New(disgord.Config{
		ProjectName: "Basic bot.go",
		BotToken:    config.Token, //Get Token From config.json
		Presence: &disgord.UpdateStatusPayload{ /* Presence that the bot will have when starting */
			Game: &disgord.Activity{
				Name: "Activity name | Disgord!",
				Type: 0,
			},
			Status: disgord.StatusIdle,
		},
	})

	/* connect, and stay connected until a system interrupt takes place */
	defer client.Gateway().StayConnectedUntilInterrupted()

	/* New router */
	router := disgordrouter.New()

	commands.NewRouter(client, router)

	client.Gateway().BotReady(func() {
		guildIDs := client.GetConnectedGuilds()
		botUser, _ := client.CurrentUser().Get()

		/*Bot#0000 | Guilds: 1*/
		log.Printf("%s | Guilds: %d", botUser.Tag(), len(guildIDs))
	})
}

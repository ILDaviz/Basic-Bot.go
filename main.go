package main //package main (main file)

import (
	"Bot/commands"
	"Bot/config"
	"context" //package context
	"log"     //package log

	"github.com/Necroforger/dgrouter/disgordrouter"
	"github.com/andersfylling/disgord" //lib disgord
)

func main() {

	client := disgord.New(disgord.Config{
		ProjectName: "Basic bot.go",
		BotToken:    config.Token, //Get Token From config.json
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "Activity name | Disgord!",
				Type: 0,
			},
			Status: disgord.StatusIdle,
		},
	})

	/* connect, and stay connected until a system interrupt takes place */
	defer client.StayConnectedUntilInterrupted(context.Background())

	router := disgordrouter.New()

	commands.NewRouter(client, router)

	client.On(disgord.EvtReady, func() {
		guilds, _ := client.GetGuilds(context.Background(), nil, disgord.IgnoreCache)
		botUser, _ := client.Myself(context.Background())

		log.Printf("%s | Guilds: %d", botUser.Tag(), len(guilds))
	})
}

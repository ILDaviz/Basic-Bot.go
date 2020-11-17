package commands

import (
	"Bot/config"
	"Bot/database"
	"context"
	"github.com/Necroforger/dgrouter/disgordrouter"
	"github.com/andersfylling/disgord"
)

type Command struct {
	Name        string
	Description string
	Category    string
	Run         disgordrouter.HandlerFunc
}

var Commands map[string]*Command = make(map[string]*Command)

var BotUser *disgord.User

func NewRouter(client *disgord.Client, router *disgordrouter.Route) {
	for _, cmd := range Commands {
		router.On(cmd.Name, cmd.Run).Desc(cmd.Description).Cat(cmd.Category)
	}

	BotUser, _ = client.Myself(context.Background())

	client.On(disgord.EvtMessageCreate, func(session disgord.Session, m *disgord.MessageCreate) {
		if m.Message.IsDirectMessage() {
			return /* Ignore DM */
		}

		router.FindAndExecute(session, getPrefix(m.Message.GuildID), BotUser.ID, m.Message)
	})
}

func NewCommand(name string, description string, category string, handler disgordrouter.HandlerFunc) *Command {
	command := &Command{
		Name:        name,
		Description: description,
		Category:    category,
		Run:         handler,
	}
	Commands[name] = command

	return command
}

func getPrefix(id disgord.Snowflake) string {

	if data, err := database.Get("SELECT prefix FROM settings WHERE GuildId = ?", id); err == nil {
		return data.(string)
	}

	_ = database.Run("INSERT INTO settings (GuildId, Prefix) VALUES(?,?)", id, config.Prefix)

	return config.Prefix
}

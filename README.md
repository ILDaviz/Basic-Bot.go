# Simple Bot.go

A simple example discord bot written in go using the [disgord](https://github.com/andersfylling/disgord) library.

## Requirements

- Have [Go](https://golang.org/) installed on your PC
- Having the disgord Library, in the terminal: `$ go get github.com/andersfylling/disgord`


### Installation

1. Clone Repository
2. Run in the terminal: `$ go build main.go` in the cloned directory
3. Execute `main.exe`

> **Note:** If you don't want to compile the bot, you can run it by running in terminal: `$ go run main.go` in the cloned directory

### Config

Fill the `config.json` file with the following template:

```json
{
    "Prefix":"PREFIX",
    "Bot_ID":"BOT-ID",
    "Bot_Token":"BOT-TOKEN"
}
```
- Prefix: The bot prefix
- Bot_ID: The bot id
- Bot_Token: the Token of the bot that is obtained in the Discord [Developers](https://discordapp.com/developers/applications) page

### Commands

> | Command | Desciption | 
> | :---------------: | :----------------: | 
> | **```ping```** | The bot responds with a pong! | 
> | **```embed```** | The bot sends an example of an embed | 
> | **```say```** | The bot sends what you say | 
> | **```avatar```** | The bot embeds the avatar of the author of the message or the mentioned one | 

### Support

If you don't know GO I do not recommend using this and I will not provide support for that.

### Author

[Night0880](https://github.com/Night0880)

**Discord:** Night#0880

Always working on a new project. Learning more every day ♥
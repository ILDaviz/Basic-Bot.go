# Basic bot.go

A simple example discord bot written in go using the [disgord](https://github.com/andersfylling/disgord) library.

## Requirements

- Have [Go](https://golang.org/) installed on your PC
- Have knowledge about the [Go](https://golang.org/) language and its [packages](https://golang.org/pkg/) and [Go Modules](https://golang.org/ref/mod)
- For windows it is necessary to have a gcc compiler for the [sqlite3 database](https://github.com/mattn/go-sqlite3#windows)


### Installation

1. Clone Repository
2. Run in the terminal: `$ go build .` in the cloned directory
3. Execute `Bot.exe` (In Linux simply `$ ./Bot`) 

> **Note:** If you don't want to compile the bot, you can run it by running in terminal: `$ go run .` in the cloned directory

### Config

Fill the `config.json` file with the following template:

```json
{
    "Prefix":"PREFIX",
    "Token":"BOT-TOKEN"
}
```
- Prefix: The bot prefix
- Token: the Token of the bot that is obtained in the Discord [Developers](https://discordapp.com/developers/applications) page

### Commands

> | Command | Desciption | Permissions |
> | :---------------: | :----------------: | :----------------: | 
> | **```ping```** | The bot responds with a pong! | NP
> | **```embed```** | The bot sends an example of an embed | NP
> | **```say```** | The bot sends what you say | NP
> | **```avatar```** | The bot embeds the avatar of the author of the message or the mentioned one | NP
> | **```kick```** | The bot kicks the mentioned user | Kick Members
> | **```ban```** | The bot ban the mentioned user | Ban Members
> | **```clear```**| Deletes a certain amount of messages on the channel in the range of `1 - 100` | Manage Messages
> | **```setprefix```**| Update the server prefix | Manage Guild
> | **```help```**| Displays the list of commands or information for a specific command | NP
> NP = No permission required
### Support

If you don't know GO I do not recommend using this and I will not provide support for that.

### Tasks to do
- ~~Sort the main file (divide it into several files)~~ 28/July/2020
- ~~Basic command handler (possibly with framework)~~ 17/November/2020

### Author

[Danny](https://github.com/Danny2105)


Always working on a new project. Learning more every day ♥
package main //package main (main file)

import (            
    "context"       //package context
    "os"            //package os
    "io/ioutil"     //package io/ioutil
    "log"           //package log
    "encoding/json" //package encoding/json
    "strings"       //package string
    "strconv"       //package strconv
   
    "github.com/andersfylling/disgord" //lib disgord
)

var (
    Client *disgord.Client; //var Client Type disgord.Client

    config = GetConfig("config/config.json"); //Get Config from JSON

    prefix = config.Prefix; //Get Prefix from config

    BotID int64

)


func main() {
	
    os.Setenv("TOKEN", config.BotToken) //SetEnv TOKEN From config
    
    Client = disgord.New(disgord.Config{
        BotToken: os.Getenv("TOKEN"), //Get Token From env
    })
    /* connect, and stay connected until a system interrupt takes place */
    defer Client.StayConnectedUntilInterrupted(context.Background())


    /*------------------------- Handler --------------------------*/
    /* create a handler and bind it to new message events
    * handlers/listener are run in sequence if you register more than one
    * so you should not need to worry about locking your objects unless you do any
    * parallel computing with said objects
    */
	Client.On(disgord.EvtReady, EventReadyHandler)

    Client.On(disgord.EvtMessageCreate, MessageCreate)

}

/*------------------- Event Ready Handler -------------------------------*/

func EventReadyHandler(){
	log.Print("Bot Ready!")
}

/*------------------------------- Event Message Handler -----------------------------*/

func MessageCreate(discord disgord.Session, m *disgord.MessageCreate){
    message := m.Message
    
	if message.Author.Bot {  //If the author of the message is the bot it returns nothing
        return
    }

    content := message.Content //Message content

    if len(content) <= len(prefix){ //If the content length is less than or equal to the prefix length, nothing returns
        return
    }

    if content[:len(prefix)] != prefix{ //If the content does not include the prefix returns nothing
        return
    }

    content = content[len(prefix):] //Content jumps over the prefix

	if len(content) < 1 { //If the content length is less than 1 it returns
		return
    }
    
    args := strings.Fields(content) //Arguments
    command := strings.ToLower(args[0]) //Command name

    
        ID, err := strconv.ParseInt(config.BotID, 0, 64); 
        if err != nil{
            log.Print("String to int64 ERROR: ", err)
            return
        }
    
      
        BotID := disgord.Snowflake(ID) //Convert bot ID to snowflake
        //BotID int
        
        BotUser, err := Client.GetUser(context.Background(), BotID, disgord.IgnoreCache) //Getting the bot as a user
        if err != nil{
            log.Print(err)
            return
        }

	if command == "ping"{ //If the command is ping

		message.Reply(context.Background(), discord, "pong!") //come back pong!
    
    }

    if command == "embed"{ //if the command is embed
	
		avatar, err := message.Author.AvatarURL(2048, true) 
		if err != nil{
			log.Print(err)
			return
        }
        /*We get the avatar of the author of the message
        <user>.AvatarURL(avatar_size, gif) 
        *avatar_size -> avatar size int
        *gif -> avatar.gif bool
        */

        avatarBot, err := BotUser.AvatarURL(2048, false)
        if err != nil{
            log.Print(err)
            return
        }
        /*we get the avatar of the bot
        */

        //A new parameter is created
		message.Reply(context.Background(), discord, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{ 
			Title: "The embed title", //Embed title
			Description: "The embed description",
			URL: "https://github.com/Night0880/", //the url of Title
			Color: 0xe9e9e9, //the embed color
			Footer: &disgord.EmbedFooter{ //Embed Footer
				Text: "Footer text", //Text of footer
				IconURL: avatarBot, //Icon of Footer
			},
			Author: &disgord.EmbedAuthor{ //Embed AUthor
				Name: message.Author.Tag(), //Text of embed Author (tag of author message)
				IconURL: avatar, //Icon of author
			},
			Fields:[]*disgord.EmbedField{ //embed fields
				{ //Field 1
				Name: "field value 1",
				Value: "field value 2",
				Inline: true,
				},
				{ //field 2
					Name: "field2 value 1",
					Value: "field2 value 2",
			
				},
			},
		},
	  })
    }
    
    if command == "say"{ //if the command is say
        if len(args) > 1{ //if args[1]
            /*args[0] -> -say
            *args[1] -> -say [arguments1]
            *args[2] -> -say [arguments1] [arguments2]
            *args[3] -> -say [arguments1] [aguments2] [arguments3]
            */
        message.Reply(context.Background(), discord, strings.Join(args[1:], " ")) 
        /*
        *Joining arguments 1 onwards
        *string.Join(args[1:], " ") Joining arguments 1 onwards and separated by spaces
        *string.Join(array, "Separation")
        */
        }else{
            /* If there were no arguments to send */
            message.Reply(context.Background(), discord, "There are no arguments to send")
        }
    }
}






/*-------------------------------- CONFIG ------------------------------*/


type Config struct{
	Prefix 		string			`json:"Prefix"`
	BotID   	string 			`json:"Bot_ID"`
	BotToken	string			`json:"Bot_Token"`
}

/* config.json
* "Prefix":"Prefix" ->          the bot prefix
* "Bot_ID":"BOT-ID" ->          The Bot ID
* "Bot_Token":"BOT-TOKEN" ->    The bot token (https://discord.com/developers/applications)
*/

func GetConfig(filename string) *Config{
    /* Read file config.json return Type Config */
	body, err := ioutil.ReadFile(filename)

	if err != nil{
		log.Print("Error loading config ", err)
		return nil
	}

	var config Config
    /* var config Type Config */
		json.Unmarshal(body, &config)

		return &config

}
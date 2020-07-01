package main //package main (main file)

import (
	"context"       //package context
	"encoding/json" //package encoding/json
	"fmt"           //package fmt
	"io/ioutil"     //package io/ioutil
	"log"           //package log
	"os"            //package os
	"strings"       //package strings
	"time"          //package time

	"github.com/andersfylling/disgord" //lib disgord
)

var (
	client *disgord.Client //var client Type disgord.Client (Global Var in this file)

	config = loadConfig("config/config.json") //Get Config from JSON (Global var in this file)

	prefix = config.Prefix //Get Prefix from config (Global var in this file)

	botUser *disgord.User //var botUser type disgord.User (Global var)
)

func main() {

	os.Setenv("TOKEN", config.BotToken) //SetEnv TOKEN From config

	client = disgord.New(disgord.Config{
		BotToken: os.Getenv("TOKEN"), //Get Token From env
	})

	/* connect, and stay connected until a system interrupt takes place */
	defer client.StayConnectedUntilInterrupted(context.Background())

	/*------------------------- Handler --------------------------*/
	/* create a handler and bind it to new message events
	 * handlers/listener are run in sequence if you register more than one
	 * so you should not need to worry about locking your objects unless you do any
	 * parallel computing with said objects
	 */
	client.On(disgord.EvtReady, eventReadyHandler)

	client.On(disgord.EvtMessageCreate, messageCreate)

}

/*------------------- Event Ready Handler -------------------------------*/

func eventReadyHandler() {

	client.UpdateStatus(&disgord.UpdateStatusPayload{
		Since: nil,
		Game: &disgord.Activity{ /*Type Activity*/
			Name: "Activity name | Disgord!",
			Type: 0,
		},
		Status: disgord.StatusIdle, /*Bot Status: StatusIdle. Also StatusOnline, StatusDnd, StatusOffline*/
		AFK:    false,
	})

	/* func(client *CLient) Myself(ctx) (*User, Error) */
	botUser, _ = client.Myself(context.Background())

	guilds, _ := client.GetGuilds(context.Background(), nil, disgord.IgnoreCache)

	log.Print(fmt.Sprintf("%s | Guilds: %d", botUser.Tag(), len(guilds)))
}

/*------------------------------- Event Message Handler -----------------------------*/

func messageCreate(discord disgord.Session, m *disgord.MessageCreate) {
	message := m.Message

	if message.Author.Bot { //If the author of the message is the bot it returns nothing
		return
	}

	content := message.Content //Message content

	if len(content) <= len(prefix) { //If the content length is less than or equal to the prefix length, nothing returns
		return
	}

	if content[:len(prefix)] != prefix { //If the content does not include the prefix returns nothing
		return
	}

	content = content[len(prefix):] //Content jumps over the prefix

	if len(content) < 1 { //If the content length is less than 1 it returns
		return
	}

	args := strings.Fields(content)     //Arguments
	command := strings.ToLower(args[0]) //Command name

	switch command {

	case "ping": //if the command is ping

		message.Reply(context.Background(), discord, "pong!") //come back pong!

		break

	case "embed": //if the command is embed

		avatar, err := message.Author.AvatarURL(2048, true)
		if err != nil {
			log.Print(err)
			return
		}
		/*We get the avatar of the author of the message
		  <user>.AvatarURL(avatar_size, gif)
		  *avatar_size -> avatar size int
		  *gif -> avatar.gif bool
		*/

		avatarBot, err := botUser.AvatarURL(2048, false)
		if err != nil {
			log.Print(err)
			return
		}
		/* we get the avatar of the bot */

		//A new parameter is created
		message.Reply(context.Background(), discord, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "The embed title", //Embed title
				Description: "The embed description",
				URL:         "https://github.com/Night0880/", //the url of Title
				Color:       0xe9e9e9,                        //the embed color int
				Timestamp: disgord.Time{
					Time: time.Now(),
				},
				Footer: &disgord.EmbedFooter{ //Embed Footer
					Text:    "Footer text", //Text of footer
					IconURL: avatarBot,     //Icon of Footer
				},
				Author: &disgord.EmbedAuthor{ //Embed AUthor
					Name:    message.Author.Tag(), //Text of embed Author (tag of author message)
					IconURL: avatar,               //Icon of author
				},
				Fields: []*disgord.EmbedField{ //embed fields
					{ //Field 1
						Name:   "field value 1",
						Value:  "field value 2",
						Inline: true,
					},
					{ //field 2
						Name:   "field2 value 1",
						Value:  "field2 value 2",
						Inline: true,
					},
					{ //field 3
						Name:   "field3 value 1",
						Value:  "field3 value 2",
						Inline: true,
					},
				},
			},
		})

		break

	case "say": //if the command is say

		if len(args) > 1 { //if args[1]
			/*args[0] -> -say -> prefixcommand
			 *args[1] -> -say [arguments1]
			 *args[2] -> -say [arguments1, arguments2]
			 *args[3] -> -say [arguments1, aguments2, arguments3]
			 */
			message.Reply(context.Background(), discord, strings.Join(args[1:], " "))
			/*
			 *Joining arguments 1 onwards
			 *string.Join(args[1:], " ") Joining arguments 1 onwards and separated by spaces
			 *string.Join(array, "Separation")
			 */
		} else {
			/* If there were no arguments to send */
			message.Reply(context.Background(), discord, "There are no arguments to send")
		}

		break

	case "avatar": //if the command is avatar

		/* Getting the mention */

		MentionUser := message.Mentions
		/*MentionUser[]*User
		 *Returns array of users mentioned
		 */

		/* Message author avatar */
		AvatarAuthor, err := message.Author.AvatarURL(2048, true)
		if err != nil {
			log.Print(err)
			return
		}

		if len(MentionUser) < 1 {
			/* If the length of the mention is less than 1, return the avatar of the author of the message */
			message.Reply(context.Background(), discord, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{ //embed
					Title: "Avatar of " + message.Author.Tag(), //Embed title
					Color: 0xe9e9e9,                            //Embed Color
					Timestamp: disgord.Time{
						Time: time.Now(),
					},
					Image: &disgord.EmbedImage{ //Embed Image
						URL: AvatarAuthor, //URL of the image in this case the avatar
					},
					Author: &disgord.EmbedAuthor{ //Embed AUthor
						Name:    message.Author.Tag(), //Text of embed Author (tag of author message)
						IconURL: AvatarAuthor,         //Icon of author
					},
				},
			})
		} else {
			/* If there were 1 or more mentions */

			AvatarMention, err := MentionUser[0].AvatarURL(2048, true)
			if err != nil {
				log.Print(err)
				return
			}

			/* Obtains the avatar of the first user mentioned */
			message.Reply(context.Background(), discord, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{ //Embed
					Title: "Avatar of " + MentionUser[0].Tag(), //Embed Title
					Color: 0xe9e9e9,                            //Embed Color
					Timestamp: disgord.Time{
						Time: time.Now(),
					},
					Image: &disgord.EmbedImage{ //EmbedImage
						URL: AvatarMention, //URL of the image in this case the avatar of the first user mentioned
					},
					Author: &disgord.EmbedAuthor{ //Embed AUthor
						Name:    message.Author.Tag(), //Text of embed Author (tag of author message)
						IconURL: AvatarAuthor,         //Icon of author
					},
				},
			})
		}

		break

	case "kick": //if the command is kick

		permissionBot, err := hasPermission(discord, disgord.PermissionKickMembers, message.GuildID, botUser.ID) //Check the "Kick Members" Permission for the bot
		if err != nil {
			log.Print(fmt.Sprintf("An error occurred while trying to verify bot permission %s", err)) //If an error occurs in verifying the permission
			return
		}

		permissionUser, err := hasPermission(discord, disgord.PermissionKickMembers, message.GuildID, message.Author.ID) //Check the "Kick members" permission for the message author
		if err != nil {
			log.Print(fmt.Sprintf("An error occurred while trying to verify user permission %s", err)) //If an error occurs in verifying the permission
		}

		if permissionUser != true {
			message.Reply(context.Background(), discord, "Sorry, but you don't have permission to kick members")
			return //If you do not have permission, you are notified and execution ends
		}

		if permissionBot != true {
			message.Reply(context.Background(), discord, "I don't have permission to kick members")
			return //If the bot does not have permission, it is notified and execution is terminated
		}

		if len(message.Mentions) < 1 {

			message.Reply(context.Background(), discord, "You must mention a member to kick")
			//If there was no mention to any user
		} else {

			if len(args) > 2 {

				err := client.KickMember(context.Background(), message.GuildID, message.Mentions[0].ID, strings.Join(args[2:], " "), disgord.IgnoreCache)
				if err != nil {
					message.Reply(context.Background(), discord, fmt.Sprintf("an unexpected error occurred: `%s`", err))
					return //If an error occurs during the kick
				}
				message.Reply(context.Background(), discord, fmt.Sprintf("`%s`, Successfully kicked!\n> **Reason**: %s", message.Mentions[0].Tag(), strings.Join(args[2:], " ")))
				//Kick executed successfully
			} else {
				//If there was no reason for the kick
				message.Reply(context.Background(), discord, "You must add a reason for the kick")
			}
		}
		break

	case "ban":

		permissionBot, err := hasPermission(discord, disgord.PermissionBanMembers, message.GuildID, botUser.ID) //Check the "Ban Members" Permission for the bot
		if err != nil {
			log.Print(fmt.Sprintf("An error occurred while trying to verify bot permission `%s`", err)) //If an error occurs in verifying the permission
			return
		}

		permissionUser, err := hasPermission(discord, disgord.PermissionBanMembers, message.GuildID, message.Author.ID) //Check the "Ban members" permission for the message author
		if err != nil {
			log.Print(fmt.Sprintf("An error occurred while trying to verify user permission `%s`", err)) //If an error occurs in verifying the permission
		}

		if permissionUser != true {
			message.Reply(context.Background(), discord, "Sorry, but you don't have permission to ban members")
			return //If you do not have permission, you are notified and execution ends
		}

		if permissionBot != true {
			message.Reply(context.Background(), discord, "I don't have permission to ban members")
			return //If the bot does not have permission, it is notified and execution is terminated
		}

		if len(message.Mentions) < 1 {
			//If there was no mention to any user
			message.Reply(context.Background(), discord, "You must mention a member to ban")

		} else {

			if len(args) > 2 {

				err := client.BanMember(context.Background(), message.GuildID, message.Mentions[0].ID, &disgord.BanMemberParams{DeleteMessageDays: 7, Reason: strings.Join(args[2:], " ")}, disgord.IgnoreCache)

				if err != nil {
					message.Reply(context.Background(), discord, fmt.Sprintf("an unexpected error occurred: `%s`", err))
					return //If an error occurs during the ban
				}

				message.Reply(context.Background(), discord, fmt.Sprintf("`%s`, Successfully banned!\n> **Reason:** %s", message.Mentions[0].Tag(), strings.Join(args[2:], " ")))
				//Ban executed successfully
			} else {
				//If there was no reason for the ban
				message.Reply(context.Background(), discord, "You must add a reason for the ban")
			}
		}

		break

	default:

		message.Reply(context.Background(), discord, "Command not found")

		break
	}

}

/*-------------------- Verify Permissions --------------------*/

func hasPermission(s disgord.Session, Perm disgord.PermissionBits, guildID disgord.Snowflake, userID disgord.Snowflake) (bool, error) {

	guild, err := s.GetGuild(context.Background(), guildID, disgord.IgnoreCache) //Get a Guild by ID
	if err != nil {
		return false, err //If an error occurs while getting the guild it returns false and an error
	}
	if guild.OwnerID == userID {
		return true, nil //Bypass the server owner
	}

	PermissionBit, err := s.GetMemberPermissions(context.Background(), guildID, userID, disgord.IgnoreCache) //Get all member permissions in bits (uint64)
	if err != nil {
		return false, err //If an error occurs while getting the Member Permissions in bits it returns false and the error
	}

	if (PermissionBit & 8) == 8 {
		return true, nil //Bypasses administrators

	} else if (PermissionBit & Perm) == Perm {
		return true, nil //If have permission it returns true
	}

	return false, nil //If don't have permission, returns false
}

/*-------------------- Config data --------------------*/

type configFile struct {
	Prefix   string `json:"Prefix"`
	BotID    string `json:"Bot_ID,omitempty"`
	BotToken string `json:"Bot_Token"`
}

/* config.json
 * "Prefix":"Prefix" ->          the bot prefix
 * "Bot_ID":"BOT-ID" ->          The Bot ID
 * "Bot_Token":"BOT-TOKEN" ->    The bot token (https://discord.com/developers/applications)
 */

func loadConfig(filename string) *configFile {
	/* Read file config.json return Type Config */
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Print("Error loading config ", err)
		return nil
	}

	var config configFile
	/* var config Type Config */
	json.Unmarshal(body, &config)

	return &config

}

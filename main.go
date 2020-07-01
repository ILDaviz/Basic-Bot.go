package main //package main (main file)

import (
	"context"       //package context
	"encoding/json" //package encoding/json
	"fmt"           //package fmt
	"io/ioutil"     //package io/ioutil
	"log"           //package log
	"os"            //package os
	"strconv"       //package strconv
	"strings"       //package strings
	"time"          //package time

	"github.com/andersfylling/disgord" //lib disgord
)

var (
	client *disgord.Client //var client Type disgord.Client (Global Var)

	config = loadConfig("config/config.json") //Get Config from JSON (Global var in this file)

	prefix = config.Prefix //Get Prefix from config (Global var in this file)

	botUser *disgord.User //var botUser type disgord.User (Global var)

	ctx = context.Background() //var ctx by context.Background() (Context) (Global var)
)

func main() {

	os.Setenv("DISCORD_TOKEN", config.BotToken) //SetEnv TOKEN From config

	client = disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_TOKEN"), //Get Token From env
	})

	/* connect, and stay connected until a system interrupt takes place */
	defer client.StayConnectedUntilInterrupted(ctx)

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
	botUser, _ = client.Myself(ctx)

	guilds, _ := client.GetGuilds(ctx, nil, disgord.IgnoreCache)

	log.Print(fmt.Sprintf("%s | Guilds: %d", botUser.Tag(), len(guilds)))
}

/*------------------------------- Event MessageCreate Handler -----------------------------*/

func messageCreate(session disgord.Session, m *disgord.MessageCreate) {
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

		message.Reply(ctx, session, "pong!") //come back pong!

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
		message.Reply(ctx, session, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "The embed title", //Embed title
				Description: "The embed description",
				URL:         "https://github.com/Night0880/", //the url of Title
				Color:       0xe9e9e9,                        //the embed color int
				Timestamp: disgord.Time{
					Time: time.Now(), //the embed timestamp
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
			message.Reply(ctx, session, strings.Join(args[1:], " "))
			/*
			 *Joining arguments 1 onwards
			 *string.Join(args[1:], " ") Joining arguments 1 onwards and separated by spaces
			 *string.Join(array, "Separation")
			 */
		} else {
			/* If there were no arguments to send */
			message.Reply(ctx, session, "There are no arguments to send")
		}

		break

	case "avatar": //if the command is avatar

		/* Getting the mention */

		MentionUser := message.Mentions
		/*MentionUser []*User
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
			message.Reply(ctx, session, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{ //embed
					Title: "Avatar of " + message.Author.Tag(), //Embed title
					Color: 0xe9e9e9,                            //Embed Color
					Timestamp: disgord.Time{
						Time: time.Now(), //embed timestamp, time.Now()(Time)
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
			message.Reply(ctx, session, &disgord.CreateMessageParams{
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

		PermsBot := hasPermission(session, botUser, message.GuildID, KICK_MEMBERS) //Check the "Kick Members" Permission for the bot

		PermsUser := hasPermission(session, message.Author, message.GuildID, KICK_MEMBERS) //Check the "Kick members" permission for the message author

		if !PermsUser {
			message.Reply(ctx, session, "Sorry, but you don't have permission to kick members")
			return //If you do not have permission, you are notified and execution ends
		}

		if !PermsBot {
			message.Reply(ctx, session, "I don't have permission to kick members")
			return //If the bot does not have permission, it is notified and execution is terminated
		}

		if len(message.Mentions) < 1 {

			message.Reply(ctx, session, "You must mention a member to kick")
			//If there was no mention to any user
		} else {

			if len(args) > 2 {

				err := client.KickMember(ctx, message.GuildID, message.Mentions[0].ID, strings.Join(args[2:], " "), disgord.IgnoreCache)
				if err != nil {
					message.Reply(ctx, session, fmt.Sprintf("an unexpected error occurred: `%s`", err))
					return //If an error occurs during the kick
				}
				message.Reply(ctx, session, fmt.Sprintf("`%s`, Successfully kicked!\n> **Reason**: %s", message.Mentions[0].Tag(), strings.Join(args[2:], " ")))
				//Kick executed successfully
			} else {
				//If there was no reason for the kick
				message.Reply(ctx, session, "You must add a reason for the kick")
			}
		}
		break

	case "ban": //if the command is ban

		PermsBot := hasPermission(session, botUser, message.GuildID, BAN_MEMBERS) //Check the "Ban Members" Permission for the bot

		PermsUser := hasPermission(session, message.Author, message.GuildID, BAN_MEMBERS) //Check the "Ban members" permission for the message author

		if !PermsUser {
			message.Reply(ctx, session, "Sorry, but you don't have permission to ban members")
			return //If you do not have permission, you are notified and execution ends
		}

		if !PermsBot {
			message.Reply(ctx, session, "I don't have permission to ban members")
			return //If the bot does not have permission, it is notified and execution is terminated
		}

		if len(message.Mentions) < 1 {
			//If there was no mention to any user
			message.Reply(ctx, session, "You must mention a member to ban")

		} else {

			if len(args) > 2 {

				err := client.BanMember(ctx, message.GuildID, message.Mentions[0].ID, &disgord.BanMemberParams{DeleteMessageDays: 7, Reason: strings.Join(args[2:], " ")}, disgord.IgnoreCache)

				if err != nil {
					message.Reply(ctx, session, fmt.Sprintf("an unexpected error occurred: `%s`", err))
					return //If an error occurs during the ban
				}

				message.Reply(ctx, session, fmt.Sprintf("`%s`, Successfully banned!\n> **Reason:** %s", message.Mentions[0].Tag(), strings.Join(args[2:], " ")))
				//Ban executed successfully
			} else {
				//If there was no reason for the ban
				message.Reply(ctx, session, "You must add a reason for the ban")
			}
		}

		break

	case "clear": //if the command is clear

		PermsUser := hasPermission(session, message.Author, message.GuildID, MANAGE_MESSAGES) //Check if the author has the permission of MANAGE_MESSAGES: Utils.go 23:6

		if PermsUser {
			//If the author has the permission, it is verified if the bot has the permission

			PermsBot := hasPermission(session, botUser, message.GuildID, MANAGE_MESSAGES) //Check if the bot has the permission of MANAGE_MESSAGES: Util.go 23:6

			if PermsBot {
				//If both have permission it is verified if there was the first argument after the command
				if len(args) > 1 {
					//If there were arguments, we convert the input to int
					number, err := strconv.Atoi(args[1])
					//args[1] <- string
					//number -> int

					if err != nil {

						message.Reply(ctx, session, "You must enter only numbers")
						//If error is different from nil, it means that the amount entered is not a number
					} else if number > 0 && number <= 100 {
						//If number is greater than 0 and less than or equal to 100 we continue
						Params := &disgord.GetMessagesParams{
							Limit: uint(number),
						}
						//We get all messages by the limit entered

						MessagesAmount, err := client.GetMessages(ctx, message.ChannelID, Params)
						//We get the number of Messages
						//MessageAmount []*disgord.Message

						if err != nil {
							message.Reply(ctx, session, fmt.Sprintf("An error occurred to verify the messages to delete\nError_ %s", err))
							//If error is different from nil it means that there was an error getting the messages
						} else {
							//Type disgord.DeleteMessagesParams
							Params := &disgord.DeleteMessagesParams{}
							for i := 0; i < len(MessagesAmount); i++ {
								//This for is in charge of adding the messages that are going to be removed from the channel
								Params.AddMessage(MessagesAmount[i])
								/*(Params *disgord.DeleteMessagesParams) AddMessage(msg *disgord.Message)*/
							}
							if err = client.DeleteMessages(ctx, message.ChannelID, Params, disgord.IgnoreCache); err != nil {
								message.Reply(ctx, session, fmt.Sprintf("They cannot be removed `%d` messages\nError: `%s`", number, err))
								return
								//If error is different from nil it means that there was an error in deleting the indicated number of messages (usually because messages of more than 2 weeks cannot be deleted)
							}
							message.Reply(ctx, session, fmt.Sprintf("I just deleted `%d` message(s)", number))
							//If everything went well, I delete the messages and it is reported that the number of messages indicated was successfully deleted.

						}
					} else {

						message.Reply(ctx, session, "The range to delete messages is **1 - 100**")
						//If you enter a negative number or greater than 100, we send this message to let you know the range that messages can be deleted.

					}

				} else {
					message.Reply(ctx, session, "You must enter the number of messages to remove")
					//If you did not put the number of messages to remove, we let you know
				}
			} else {
				message.Reply(ctx, session, "I don't have the `Manage messages` permission")
				//If the bot does not have permissions to manage messages, we let you know that you need it.
			}

		} else {
			message.Reply(ctx, session, "You do not have the `Manage messages` permission")
			//If the author does not have permission to manage messages, we let him know that he cannot execute this command without that permission.

		}
		break

	default:

		message.Reply(ctx, session, "Command not found")
		//If the command entered is not registered, we will let you know
		break
	}

}

/*--------------- List ALL Permissions -------------------*/
var (
	CREATE_INSTANT_INVITE = disgord.PermissionCreateInstantInvite
	KICK_MEMBERS          = disgord.PermissionKickMembers
	BAN_MEMBERS           = disgord.PermissionBanMembers
	ADMINISTRATOR         = disgord.PermissionAdministrator
	MANAGE_CHANNELS       = disgord.PermissionManageChannels
	MANAGE_GUILD          = disgord.PermissionManageServer
	ADD_REACTIONS         = disgord.PermissionAddReactions
	VIEW_AUDIT_LOG        = disgord.PermissionViewAuditLogs
	VIEW_CHANNEL          = disgord.PermissionReadMessages
	SEND_MESSAGES         = disgord.PermissionSendMessages
	SEND_TTS_MESSAGES     = disgord.PermissionSendTTSMessages
	MANAGE_MESSAGES       = disgord.PermissionManageMessages
	EMBED_LINKS           = disgord.PermissionEmbedLinks
	ATTACH_FILES          = disgord.PermissionAttachFiles
	READ_MESSAGE_HISTORY  = disgord.PermissionReadMessageHistory
	MENTION_EVERYONE      = disgord.PermissionMentionEveryone
	USE_EXTERNAL_EMOJIS   = disgord.PermissionUseExternalEmojis
	CONNECT               = disgord.PermissionVoiceConnect
	SPEAK                 = disgord.PermissionVoiceSpeak
	MUTE_MEMBERS          = disgord.PermissionVoiceMuteMembers
	DEAFEN_MEMBERS        = disgord.PermissionVoiceDeafenMembers
	MOVE_MEMBERS          = disgord.PermissionVoiceMoveMembers
	USE_VAD               = disgord.PermissionVoiceUseVAD
	PRIORITY_SPEAKER      = disgord.PermissionVoicePrioritySpeaker
	CHANGE_NICKNAME       = disgord.PermissionChangeNickname
	MANAGE_NICKNAMES      = disgord.PermissionManageNicknames
	MANAGE_ROLES          = disgord.PermissionManageRoles
	MANAGE_WEBHOOKS       = disgord.PermissionManageWebhooks
	MANAGE_EMOJIS         = disgord.PermissionManageEmojis
)

/*-------------------- Verify Permissions --------------------*/
func hasPermission(session disgord.Session, User *disgord.User, GuildID disgord.Snowflake, Permission uint64) bool {
	Member, err := session.GetMember(ctx, GuildID, User.ID, disgord.IgnoreCache)
	if err != nil {
		//If an error occurs when trying to get the user as a member we return false and the error to the console
		log.Print(err)
		return false
	}
	for _, roleID := range Member.Roles {
		guild, err := session.GetGuild(ctx, GuildID, disgord.IgnoreCache)
		if err != nil {
			//If an error occurs when trying to get the server we return false and the error to the console
			log.Print(err)
			return false
		}

		if guild.OwnerID == Member.UserID {
			return true //If the user ID is equal to the owner's id, we omit the owner of the server (since it has all the permissions)
		}

		role, err := guild.Role(roleID)
		if err != nil {
			//If an error occurs when trying to obtain the Role, we return false and the error to the console
			log.Print(err)
			return false
		}
		if (role.Permissions & 0x8) == 0x8 {
			return true //If you have administrator permission we omit it (since with administrator permission it is to have all the permissions)

		}

		if (role.Permissions & Permission) == Permission {
			//If you have the indicated permission we return true
			return true
		}
	}
	return false //If you do not own the server, you do not have administrator permission and you do not have the permission to search, then we return false
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
	/*   config Type Config */
	json.Unmarshal(body, &config)

	return &config

}

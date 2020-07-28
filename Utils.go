package main

import (
	"io/ioutil"     //package io/ioutil
	"encoding/json" //package encoding/json
	"log"           //package log

	"github.com/andersfylling/disgord" //lib disgord
)


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
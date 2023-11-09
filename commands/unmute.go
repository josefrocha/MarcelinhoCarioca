package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Unmute(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	guildID := message.GuildID
	memberID := args[1]

	if memberID == "" {
		client.ChannelMessageSend(guildID, "Você não inseriu o id do usuário.")
		return
	}

	if (len(message.Mentions) > 0) {
		client.ChannelMessageSend(guildID, "Não há como desmutar com menção, apenas por ID.")
		return
	}

	member, err := client.GuildMember(message.GuildID, memberID)
	userID := member.User.ID
	if err != nil {
		client.ChannelMessageSend(message.ChannelID, "Este usuário não está no servidor.")
		return
	}

	err = client.GuildMemberTimeout(guildID, userID, nil)
	if (err != nil) {
		client.ChannelMessageSend(message.ChannelID, "Esse usuário não está mutado ou não tenho permissão suficiente.")
		return
	}
	
	client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v foi desmutado", member.User.Username))
}

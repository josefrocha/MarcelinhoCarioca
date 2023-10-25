package commands 

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Unban(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	if (len(args) < 2) {
		client.ChannelMessageSend(message.ChannelID, "Você não passou o ID.")
		return
	}

	guildID := message.GuildID
    userArgumentId := args[1]

	if (userArgumentId == "") {
        client.ChannelMessageSend(guildID, "Você não inseriu o id do usuário.")
        return	
	}

	user, err := client.User(userArgumentId);
    if (err != nil) {
	    client.ChannelMessageSend(guildID, "Esse usuário não está no servidor ou o ID é invalido.")
		return
	}

	username := user.Username
	userID := user.ID 

	err = client.GuildBanDelete(guildID, userID);
	if (err != nil) {
		client.ChannelMessageSend(message.ChannelID, "Esse usuário não está banido ou não tenho permissão suficiente.")
		return
	}

	client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v foi desbanido.", username))
}
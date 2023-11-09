package commands 

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Unban(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	guildID := message.GuildID
	if (len(args) < 2) {
		client.ChannelMessageSend(message.ChannelID, "Você não passou o ID.")
		return
	}
	
	if (len(message.Mentions) > 0) {
		client.ChannelMessageSend(guildID, "Não há como banir com menção, apenas por ID.")
		return
	}
	
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
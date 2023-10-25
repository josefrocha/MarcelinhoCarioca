package commands 

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

func Ban(client *discordgo.Session, message *discordgo.MessageCreate, args) {
	guildID := message.GuildID
    memberID := args[1]
	reason := args[2]

	if (memberID == "") {
        client.ChannelMessageSend(guildID, "Você não inseriu o id do usuário.")
	}

	if (reason == "") {
		reason = "Sem motivo"
	}

	member, err := client.GuildMember(guildID, memberID);
	user := member.User
    if (err != nil) {
	    client.ChannelMessageSend(guildID, "Esse usuário não está no servidor ou o ID é invalido.")
	} 

	err := client.GuildBanCreateWithReason(guildID, user, reason, 1);
	if (err != nil) {
		fmt.Println(err)
	}
}
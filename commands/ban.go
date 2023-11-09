package commands 

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Ban(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	if (len(args) < 2) {
		client.ChannelMessageSend(message.ChannelID, "Você não passou o ID nem o motivo.")
		return
	}

	guildID := message.GuildID
    memberID := args[1]
	reason := "Sem motivo";

	if (len(message.Mentions) > 0) {
		client.ChannelMessageSend(guildID, "Não há como banir por menção, apenas por ID.")
		return
	}

	if (memberID == "") {
        client.ChannelMessageSend(guildID, "Você não inseriu o id do usuário.")
        return	
	}

	if memberID == message.Author.ID {
		client.ChannelMessageSend(guildID, "Vai se banir? tem certeza?!")
		return
	}

	if (len(args) > 1) {
		reasonText := args[2:]
		reason = strings.Join(reasonText, " ")
	}

	member, err := client.GuildMember(guildID, memberID);
    if (err != nil) {
	    client.ChannelMessageSend(guildID, "Esse usuário não está no servidor ou o ID é invalido.")
		return
	}

	username := member.User.Username
	userID := member.User.ID 

	err = client.GuildBanCreateWithReason(guildID, userID, reason, 1);
	if (err != nil) {
		client.ChannelMessageSend(message.ChannelID, "Não foi possivel banir o membro, eu não tenho permissão suficiente.")
		return
	}

	client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("**%v** foi banido.\nMotivo: `%v`", username, reason))
}
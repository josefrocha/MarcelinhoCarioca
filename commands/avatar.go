package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Avatar(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	var memberID string;
	
	if (len(args) < 2) {
		memberID = message.Author.ID
	} else if(len(args) == 2) {
        memberID = args[1] 
	}
	
	member, err := client.GuildMember(message.GuildID, memberID)
	if err != nil {
		client.ChannelMessageSend(message.ChannelID, "Este usuário não está no servidor.");
		return
	}

	avatar := member.AvatarURL("4096")

	imageEmbed := &discordgo.MessageEmbedImage{ URL: avatar }
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Avatar do %v (%v)", member.User.Username, member.User.ID),
		Description: fmt.Sprintf("[Download](%v)", avatar),
		Image: imageEmbed,
		Color: 0x2b2d31,
	}

	client.ChannelMessageSendEmbed(message.ChannelID, embed)
}
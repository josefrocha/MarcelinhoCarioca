package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Avatar(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	memberID := args[1]
	if memberID == "" {
		memberID = message.Author.ID
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
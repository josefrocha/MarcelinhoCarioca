package commands

import (
	"fmt"
	"time"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Mute(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	guildID := message.GuildID

	if len(args) < 3 {
		client.ChannelMessageSend(guildID, "Você não inseriu o id do usuário e nem o tempo do mute (em segundos).")
		return
	}

	var timeout *time.Time
	memberID := args[1]
	timeu, _ := strconv.Atoi(args[2])
	
	if args[2] == "" {
		client.ChannelMessageSend(guildID, "Insira o tempo. (em segundos)")
		return
	} else if timeu < 60 {
		client.ChannelMessageSend(guildID, "O mínimo é 60 segundos.")
		return
	} else {
		timeoutTime := secondsToTime(int64(timeu))
		timeout = &timeoutTime
	}

	if memberID == "" {
		client.ChannelMessageSend(guildID, "Você não inseriu o id do usuário.")
		return
	}

	if memberID == message.Author.ID {
		client.ChannelMessageSend(guildID, "Vai se mutar? tem certeza?!")
		return
	}

	member, err := client.GuildMember(message.GuildID, memberID)
	userID := member.User.ID
	if err != nil {
		client.ChannelMessageSend(message.ChannelID, "Este usuário não está no servidor.")
		return
	}

	err = client.GuildMemberTimeout(guildID, userID, timeout)
	if (err != nil) {
		client.ChannelMessageSend(message.ChannelID, "Esse usuário já está mutado ou não tenho permissão suficiente.")
		return
	}
	client.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v foi mutado por %v segundos (%v horas)", member.User.Username, args[2], timeout.Hour() - time.Now().Hour()))
}

func secondsToTime(seconds int64) time.Time {
	return time.Now().Add(time.Second * time.Duration(seconds))
}
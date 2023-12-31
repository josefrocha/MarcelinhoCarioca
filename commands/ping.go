package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(client *discordgo.Session, message *discordgo.MessageCreate,  args []string) {
	latency := client.HeartbeatLatency().Milliseconds();	

	client.ChannelMessageSend(message.ChannelID, fmt.Sprint(latency) + "ms")
}
package event

import (
	"github.com/bwmarrin/discordgo"

	"marcelinhocarioca/events"
)

func Handler(client *discordgo.Session) {
    client.AddHandler(events.MessageCreate);
	client.AddHandler(events.Ready);
}
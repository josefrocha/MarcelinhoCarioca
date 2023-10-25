package events

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

func Ready(client *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("bot is ready")
}

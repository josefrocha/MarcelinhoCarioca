package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"marcelinhocarioca/handlers/command"
)

func MessageCreate(client *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == client.State.User.ID {
		return
	}

	prefix := "?"

	if !strings.HasPrefix(message.Content, prefix) {
		return
	}

	content := strings.TrimPrefix(message.Content, prefix)
	args := strings.Fields(content)

	if len(args) > 0 {
		commands := command.Cmds()
		commandName := args[0]

		for _, cmd := range commands {
			if cmd.Name() == commandName {
				cmd.Run(client, message, args)
				break;
			}
		}
	}
}

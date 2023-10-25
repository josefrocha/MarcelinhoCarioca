package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Name() string
	Description() string
}

type Cmd struct {
	name        string
	description string
}

func (c Cmd) Name() string {
	return c.name
}

func (c Cmd) Description() string {
	return c.description
}

func Cmds() []Command {
	var cmds = []Command{
		Cmd {
			name:        "avatar",
			description: "Mostra o seu avatar ou de outro usuário específico.",
		},
		Cmd {
			name:        "help",
			description: "Lista os comandos.",
		},
		Cmd {
			name:        "ping",
			description: "Mostra o ping do websocket do bot.",
		},
		Cmd{
			name:        "ban",
			description: "Bane o usuário do servidor.",
		},
		Cmd{
			name:        "unban",
			description: "Desbane o usuário do servidor.",
		},
	}

	return cmds
}

func Help(client *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	fields := []*discordgo.MessageEmbedField{};

	commands := Cmds();

	for _, command := range commands {
		field := &discordgo.MessageEmbedField{
			Name: command.Name(),
			Value: command.Description(),
			Inline: false,
		}

		fields = append(fields, field);
	}

	embed := &discordgo.MessageEmbed{
       Title: "Lista de comandos",
	   Description: "Aqui estão os comandos disponíveis:",
	   Fields: fields,
	   Color: 0x2b2d31,
	}

	client.ChannelMessageSendEmbed(message.ChannelID, embed)
}
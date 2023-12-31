package command

import (
	"github.com/bwmarrin/discordgo"

	"marcelinhocarioca/commands"
)

type Command interface {
	Name() string
	Description() string
	Run(client *discordgo.Session, message *discordgo.MessageCreate,  args []string)
}

type Cmd struct {
	name        string
	description string
	run         func(client *discordgo.Session, message *discordgo.MessageCreate,  args []string)
}

func (c Cmd) Name() string {
	return c.name
}

func (c Cmd) Description() string {
	return c.description
}

func (c Cmd) Run(client *discordgo.Session, message *discordgo.MessageCreate,  args []string) {
	c.run(client, message, args)
}

func Cmds() []Command {
	var cmds = []Command{
		Cmd{
			name:        "avatar",
			description: "Mostra o seu avatar ou de outro usuário específico.",
			run:         commands.Avatar,
		},
		Cmd{
			name:        "help",
			description: "Lista os comandos.",
			run:         commands.Help,
		},
		Cmd{
			name:        "ping",
			description: "Mostra o ping do websocket do bot.",
			run:         commands.Ping,
		},
		Cmd{
			name:        "ban",
			description: "Bane o usuário do servidor.",
			run:         commands.Ban,
		},
		Cmd{
			name:        "unban",
			description: "Desbane o usuário do servidor.",
			run:         commands.Unban,
		},
		Cmd{
			name:        "mute",
			description: "Muta o usuário por um tempo determinado no servidor.",
			run:         commands.Mute,
		},
		Cmd{
			name:        "unmute",
			description: "Muta o usuário por um tempo determinado no servidor.",
			run:         commands.Unmute,
		},
	}

	return cmds
}
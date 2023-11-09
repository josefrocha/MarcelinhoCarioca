package main

import (
	"os"
	"os/signal"
	"log"
	
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"marcelinhocarioca/handlers/event"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	TOKEN_BOT := os.Getenv("TOKEN_BOT")

	run(TOKEN_BOT)
}

func run(TOKEN_BOT string) {
	client, err := discordgo.New("Bot " + TOKEN_BOT)

	if err != nil {
		log.Fatal(err)
	}

    event.Handler(client)

	client.Identify.Intents = discordgo.IntentsGuildMessages

	client.Open()
	defer client.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

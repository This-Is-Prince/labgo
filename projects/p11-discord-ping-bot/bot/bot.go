package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var bot = struct {
	Id      string
	Session *discordgo.Session
}{}

func Start() {
	token := os.Getenv("TOKEN")
	var err error

	bot.Session, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error: creating new discord session: %v", err.Error())
		return
	}

	user, err := bot.Session.User("@me")
	if err != nil {
		log.Fatalf("Error: getting current user: %v", err.Error())
		return
	}

	bot.Id = user.ID

	bot.Session.AddHandler(messageHandler)

	err = bot.Session.Open()
	if err != nil {
		log.Fatalf("Error: creating websocket connection to discord: %v", err.Error())
		return
	}

	fmt.Println("Bot is running!!!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == bot.Id {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}

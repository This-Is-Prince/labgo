package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	fmt.Println("Go Slack Test Bot")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}
	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")

	bot := slacker.NewClient(slackBotToken, slackAppToken)

	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Description: "ping pong chat cmd",
		Examples:    []string{"ping"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}

}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("---------Command Events---------")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

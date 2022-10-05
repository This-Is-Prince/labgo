package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	fmt.Println("Go Slack File Bot")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	api := slack.New(slackBotToken)
	channels := []string{channelId}
	files := []string{"main.go"}

	for _, file := range files {
		params := slack.FileUploadParameters{
			Channels: channels,
			File:     file,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
}

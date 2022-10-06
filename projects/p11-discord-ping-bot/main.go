package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/This-Is-Prince/discord-ping-bot/bot"
	"github.com/joho/godotenv"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Go Discord Ping Bot")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	wg.Add(1)
	bot.Start()
	wg.Wait()
}

package main

import (
	"github.com/tronget/weather-app-bot/botutil"
	"log"
)

func main() {
	bot, err := botutil.Init()
	if err != nil {
		log.Panicf("Error initializing bot: %v", err)
	}

	bot.Debug = true

	updateConfig := botutil.GetUpdate(0, 30)
	botutil.HandleMessages(bot, updateConfig)
}

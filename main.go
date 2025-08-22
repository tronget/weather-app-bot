package main

import (
	"github.com/tronget/weather-app-bot/botutil"
	"github.com/tronget/weather-app-bot/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Panicf("Error loading config: %v", err)
	}

	bot, err := botutil.Init(cfg)
	if err != nil {
		log.Panicf("Error initializing bot: %v", err)
	}

	bot.Debug = true

	updateConfig := botutil.GetUpdate(0, 30)
	botutil.HandleMessages(bot, updateConfig, cfg)
}

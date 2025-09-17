package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tronget/weather-app-bot/internal/botutil"
	"github.com/tronget/weather-app-bot/internal/config"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Panicf("Error loading config: %v", err)
	}

	bot, err := botutil.Init(cfg)
	if err != nil {
		log.Panicf("Error initializing bot: %v", err)
	}

	//bot.Debug = true

	updateConfig := botutil.GetUpdate(0, 30)
	botutil.HandleMessages(bot, updateConfig, cfg)
}

package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tronget/weather-app-bot/internal/botutil"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/network/db"
	"log"
)

func main() {
	// close the global database connection when the main function ends
	defer db.CloseDatabase()

	if err := godotenv.Load(); err != nil {
		log.Panicf("No .env file found")
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

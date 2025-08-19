package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
)

func InitBot() (*tgbotapi.BotAPI, error) {
	token := config.GetToken()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return bot, nil
}

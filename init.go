package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"sync"
)

var mutex sync.Mutex

var botInstance *tgbotapi.BotAPI

func GetInstance() (*tgbotapi.BotAPI, error) {
	if botInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if botInstance == nil {
			var err error
			botInstance, err = InitBot()
			if err != nil {
				return nil, err
			}
		}
	}
	return botInstance, nil
}

func InitBot() (*tgbotapi.BotAPI, error) {
	token, err := config.GetTelegramToken()
	if err != nil {
		return nil, err
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return bot, nil
}

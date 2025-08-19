package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"log"
)

func main() {
	bot, err := InitBot()
	if err != nil {
		log.Panicf("Error initializing bot: %v", err)
	}
	handleMessages(bot)
}

func handleMessages(bot *tgbotapi.BotAPI) {
	updateConfig := config.GetUpdate()
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			fmt.Println("this is just keep alive")
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You typed: "+update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

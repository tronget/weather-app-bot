package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/weather"
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
			continue
		}

		chatID := update.Message.Chat.ID
		messageText := update.Message.Text
		if messageText == "" {
			replyMessageText := "PLS, send me a text message with the name of the place bro.."
			msg := tgbotapi.NewMessage(chatID, replyMessageText)
			sendMessage(&msg, &update, bot)
			continue
		}

		cities, err := weather.GetCity(messageText)
		if err != nil {
			replyMessageText := "Error occurred during request. Sorry, maybe we have some problems."
			msg := tgbotapi.NewMessage(chatID, replyMessageText)
			sendMessage(&msg, &update, bot)
			continue
		}

		if len(cities) == 0 {
			replyMessageText := fmt.Sprintf("City \"%s\" not found.", messageText)
			msg := tgbotapi.NewMessage(chatID, replyMessageText)
			sendMessage(&msg, &update, bot)
			continue
		}

		city := cities[0]
		replyMessageText := fmt.Sprintf("%s\n%f : %f, ", city.Name, city.Lon, city.Lat)
		msg := tgbotapi.NewMessage(chatID, replyMessageText)
		sendMessage(&msg, &update, bot)
	}
}

func sendMessage(msg *tgbotapi.MessageConfig, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg.ReplyToMessageID = update.Message.MessageID
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error: sending message to user @%s: %v\n", update.Message.From.UserName, err)
	}
}

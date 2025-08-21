package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/weather"
	"log"
	"time"
)

func main() {
	//bot, err := InitBot()
	//if err != nil {
	//	log.Panicf("Error initializing bot: %v", err)
	//}
	//handleMessages(bot)
	go func() {
		_, err := GetInstance()
		if err != nil {
			panic(err)
		}
	}()
	_, err := GetInstance()
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
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
		msg := tgbotapi.NewMessage(chatID, "")
		if messageText == "" {
			msg.Text = "PLS, send me a text message with the name of the place bro.."
			SendMessage(&msg, &update, bot)
			continue
		}

		if update.Message.IsCommand() {
			HandleCommand()
		}

		cities, err := weather.GetCity(messageText)
		if err != nil {
			replyMessageText := "Error occurred during request. Sorry, maybe we have some problems. %v"
			msg.Text = fmt.Sprintf(replyMessageText, err)
			SendMessage(&msg, &update, bot)
			continue
		}

		if len(cities) == 0 {
			replyMessageText := fmt.Sprintf("City \"%s\" not found.", messageText)
			msg.Text = replyMessageText
			SendMessage(&msg, &update, bot)
			continue
		}

		city := cities[0]
		replyMessageText := fmt.Sprintf("%s\n%f : %f, ", city.Name, city.Lon, city.Lat)
		msg.Text = replyMessageText
		SendMessage(&msg, &update, bot)
	}
}

func SendMessage(msg *tgbotapi.MessageConfig, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg.ReplyToMessageID = update.Message.MessageID
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error: sending message to user @%s: %v\n", update.Message.From.UserName, err)
	}
}

func HandleCommand() {

}

package botutil

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/commands"
	"github.com/tronget/weather-app-bot/config"
	"log"
)

func Init(cfg *config.Config) (*tgbotapi.BotAPI, error) {
	token := cfg.TelegramToken()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("creating new bot api: %w", err)
	}
	return bot, nil
}

func GetReplyMessage(update *tgbotapi.Update, cfg *config.Config) string {
	messageText := update.Message.Text
	var replyMessageText string

	switch {
	case messageText == "":
		replyMessageText = "PLS, send me a text message with the name of the place bro.."
	case update.Message.IsCommand():
		command := update.Message.Command()
		replyMessageText = commands.Handle(command)
	default:
		replyMessageText = commands.HandleDefault(messageText, cfg)
	}

	return replyMessageText
}

func SendMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update *tgbotapi.Update) {
	msg.ReplyToMessageID = update.Message.MessageID
	if _, err := bot.Send(msg); err != nil {
		username := update.Message.From.UserName
		log.Printf("Error: sending message to user @%s: %v\n", username, err)
	}
}

func HandleMessages(bot *tgbotapi.BotAPI, updateConfig tgbotapi.UpdateConfig, cfg *config.Config) {
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatID, "")

		msg.Text = GetReplyMessage(&update, cfg)

		SendMessage(bot, &msg, &update)
	}
}

func GetUpdate(offset, timeout int) tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(offset)
	updateConfig.Timeout = timeout
	return updateConfig
}

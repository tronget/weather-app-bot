package botutil

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/commands"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/locales"
	"log"
	"strings"
)

func Init(cfg *config.Config) (*tgbotapi.BotAPI, error) {
	token := cfg.TelegramToken()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("creating new bot api: %w", err)
	}
	return bot, nil
}

func HandleMsg(cfg *config.Config, update *tgbotapi.Update, msgConfig *tgbotapi.MessageConfig) {
	messageText := update.Message.Text

	userLang := cfg.UserLanguage(update.Message.From.ID)
	fmt.Println(userLang)
	if userLang == "" {
		userLang = update.Message.From.LanguageCode
		cfg.SetUserLanguage(update.Message.From.ID, userLang)
	}

	var replyMessageText string

	switch {
	case messageText == "":
		replyMessageText = locales.Translate(locales.EMPTY_MESSAGE, userLang)
	case update.Message.IsCommand():
		commandName := update.Message.Command()
		replyMessageText = commands.Handle(commandName, msgConfig, userLang)
	default:
		replyMessageText = commands.HandleDefault(update, cfg, userLang)
	}

	msgConfig.Text = replyMessageText
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
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			msgConfig := tgbotapi.NewMessage(chatID, "")

			HandleMsg(cfg, &update, &msgConfig)

			SendMessage(bot, &msgConfig, &update)

		} else if callback := update.CallbackQuery; callback != nil {
			HandleCallback(bot, callback, cfg)
		}
	}
}

func HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, cfg *config.Config) {
	if strings.HasPrefix(callback.Data, "lang_") {
		lang := strings.TrimPrefix(callback.Data, "lang_")

		// TODO: save lang in database, not in-memory
		cfg.SetUserLanguage(callback.From.ID, lang)

		// Respond to the callback query, telling Telegram to show the user
		// a message with the data received.
		callbackMsg := locales.Translate(locales.LANG_CHOSEN, lang)
		newCallback := tgbotapi.NewCallback(callback.ID, callbackMsg)
		if _, err := bot.Request(newCallback); err != nil {
			log.Printf("accepting callback: %v", err)
		}

		// text for edited message from bot
		formatString := locales.Translate(locales.LANG_SAVED, lang)
		text := fmt.Sprintf(formatString, lang)

		chatID := callback.Message.Chat.ID
		editMessage := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, text)
		if _, err := bot.Send(editMessage); err != nil {
			log.Println("sending edited message:", err)
		}
	}
}

func GetUpdate(offset, timeout int) tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(offset)
	updateConfig.Timeout = timeout
	return updateConfig
}

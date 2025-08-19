package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"log"
	"time"
)

type Data struct {
	Info      string    `json:"info"`
	Arr       []int     `json:"arr"`
	IsTrue    bool      `json:"is_true"`
	EmbedData EmbedData `json:"zzz"`
}

type EmbedData struct {
	Info string
}

type RequestData struct {
	Key        string `json:"key"`
	AnotherKey int    `json:"another_key"`
}

func main() {
	bot, err := InitBot()
	if err != nil {
		log.Panicf("Error initializing bot: %v", err)
	}
	handleMessages(bot)
	//data := Data{
	//	Info:   "i want to fuck you",
	//	Arr:    []int{228, 1337, 333},
	//	IsTrue: true,
	//	EmbedData: EmbedData{
	//		Info: "sassaaaaaay",
	//	},
	//}
	//http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
	//
	//	switch r.Method {
	//	case http.MethodGet:
	//		jsn, err := json.Marshal(data)
	//		if err != nil {
	//			msg := "Internal server error: unable to marshal json"
	//			log.Println(msg)
	//			http.Error(w, msg, http.StatusInternalServerError)
	//			return
	//		}
	//		w.Write(jsn)
	//	case http.MethodPost:
	//		bodyBytes, err := io.ReadAll(r.Body)
	//		if err != nil {
	//			log.Println(err)
	//			http.Error(w, "Error reading request body", http.StatusBadRequest)
	//		}
	//		defer r.Body.Close()
	//
	//		rd := RequestData{}
	//		if err := json.Unmarshal(bodyBytes, &rd); err != nil {
	//
	//		}
	//		log.Println(rd)
	//	default:
	//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//	}
	//})
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessages(bot *tgbotapi.BotAPI) {
	updateConfig := config.GetUpdate()
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		time.Sleep(time.Second)
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			fmt.Println("this is just keep alive")
			continue
		}

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
}

package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		isUrl := tgbotapi.MessageEntity.IsURL(tgbotapi.MessageEntity{Type: "url"})
		if !isUrl {
			// ignore any message not a URL
			continue
		}
		if !strings.Contains(update.Message.Text, "twitter.com") && !strings.Contains(update.Message.Text, "instagram.com") {
			// ignore any message not containing twitter or instagram link
			continue
		}

		var newURL string
		newURL = strings.Replace(update.Message.Text, "twitter.com", "vxtwitter.com", 1)
		newURL = strings.Replace(newURL, "instagram.com", "ddinstagram.com", 1)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		msg.Text = newURL

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

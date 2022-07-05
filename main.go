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

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		isUrl := tgbotapi.MessageEntity.IsURL(tgbotapi.MessageEntity{Type: "url"})
		isCommand := update.Message.IsCommand()

		if !isUrl && !isCommand {
			// ignore any message not a URL or command
			continue
		}

		switch update.Message.Command() {
		case "start":
			msg.Text = "The bot is intended to replace twitter and instagram links with embedeable versions\nBuilt by @anirudh109"
			continue
		case "source":
			msg.Text = "The source code is on https://github.com/anirudhgupta109/vxgotgbot"
			continue
		case "help":
			msg.Text = "Just add to your groups and it'll automatically send a vxtwitter for twitter and ddinstagram for instagram link\nKnown commands are /source and /start"
			continue
		}

		if !strings.Contains(update.Message.Text, "twitter.com") && !strings.Contains(update.Message.Text, "instagram.com") {
			// ignore any message not containing twitter or instagram link
			continue
		}

		var newURL string
		newURL = strings.Replace(update.Message.Text, "twitter.com", "vxtwitter.com", 1)
		newURL = strings.Replace(newURL, "instagram.com", "ddinstagram.com", 1)

		msg.Text = newURL

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

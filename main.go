package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var text string
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message updates
			continue
		}

		isUrl := tgbotapi.MessageEntity.IsURL(tgbotapi.MessageEntity{Type: "url"})
		isCommand := update.Message.IsCommand()

		if !isUrl && !isCommand {
			// ignore any message not a URL or command
			continue
		}

		switch update.Message.Command() {
		case "start":
			text = "The bot is intended to replace twitter and instagram links with embedeable versions\nBuilt by @anirudh109"
			sendMessage(text, update, bot)
			continue
		case "source":
			text = "The source code is on https://github.com/anirudhgupta109/vxgotgbot"
			sendMessage(text, update, bot)

			continue
		case "help":
			text = "Just add to your groups and it'll automatically send a vxtwitter for twitter and ddinstagram for instagram link\nKnown commands are /source and /start"
			sendMessage(text, update, bot)
			continue
		}

		messageText := update.Message.Text

		if strings.Contains(messageText, "https://twitter.com/") {
			messageText = strings.Replace(messageText, "https://twitter.com/", "https://vxtwitter.com/", 1)
		}

		if strings.Contains(messageText, "https://instagram.com/") {
			messageText = strings.Replace(messageText, "https://instagram.com/", "https://ddinstagram.com/", 1)
		}

		sendMessage(messageText, update, bot)
	}
}

func sendMessage(text string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = text

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

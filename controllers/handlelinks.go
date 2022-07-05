package controllers

import (
	"strings"
	"tgbot/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleLinks(messageText string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if strings.Contains(messageText, "https://twitter.com/") {
		messageText = strings.Replace(messageText, "https://twitter.com/", "https://vxtwitter.com/", 1)
		helpers.SendMessage(messageText, update, bot)
		return
	}

	if strings.Contains(messageText, "https://instagram.com/") {
		messageText = strings.Replace(messageText, "https://instagram.com/", "https://ddinstagram.com/", 1)
		helpers.SendMessage(messageText, update, bot)
		return
	}

}

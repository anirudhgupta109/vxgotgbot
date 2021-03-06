package controllers

import (
	"strings"
	"tgbot/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
 * Function to handle substring replace of twitter and instagram links for proper embedded information
 **/
func HandleLinks(messageText string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if strings.Contains(messageText, "https://twitter.com/") {
		messageText = strings.Replace(messageText, "https://twitter.com/", "https://vxtwitter.com/", 1)
		helpers.SendMessage(messageText, update, bot)
		return
	}

	if strings.Contains(messageText, "https://www.instagram.com/") {
		messageText = strings.Replace(messageText, "https://www.instagram.com/", "https://www.ddinstagram.com/", 1)
		helpers.SendMessage(messageText, update, bot)
		return
	}

}

package controllers

import (
	"strings"
	"tgbot/controllers/socials"
	"tgbot/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
 * Function to handle substring replace of twitter and instagram links for proper embedded information
 **/
func HandleLinks(message *tgbotapi.Message, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	messageText := helpers.StripURLs(message)
	for _, domain := range helpers.GetDomains() {
		if strings.Contains(messageText, domain) {
			messageText = socials.HandleTwitter(messageText)
			messageText = socials.HandleInstagram(messageText)
			helpers.SendMessage(messageText, update, bot)
			return
		}
	}
}

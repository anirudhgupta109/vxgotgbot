package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message updates
			continue
		}

		isUrl := tgbotapi.MessageEntity.IsURL(tgbotapi.MessageEntity{Type: "url"})
		isCommand := update.Message.IsCommand()

		if isCommand {
			HandleCommands(update, bot)
			continue
		} else if isUrl {
			HandleLinks(update.Message.Text, update, bot)
		} else {
			continue
		}
	}
}

package controllers

import (
	"tgbot/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
 * Function to handle commands and execute only on a match of preset bot commands
 **/
func HandleCommands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var text string

	switch update.Message.Command() {
	case "start":
		text =
			"The bot is intended to replace twitter and instagram links with embedeable versions\nBuilt by @anirudh109"
		helpers.SendMessage(text, update, bot)
		return

	case "source":
		text =
			"The source code is on https://github.com/anirudhgupta109/vxgotgbot"
		helpers.SendMessage(text, update, bot)
		return

	case "help":
		text =
			"Just add to your groups and it'll automatically send a vxtwitter for twitter and ddinstagram for instagram link\nKnown commands are /source and /start"
		helpers.SendMessage(text, update, bot)
		return
	}
}

package helpers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(text string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = text

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

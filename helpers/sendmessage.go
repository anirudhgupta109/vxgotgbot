package helpers

import (
	"log"
	"net/url"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/**
 * Function to send a message via the bot to telegram
 **/
func SendMessage(text string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = text

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func StripURLs(message *tgbotapi.Message) string {
	if message.Text == "" {
		return ""
	}

	urls := getURLsFromMessage(message)
	finalText := message.Text

	for _, u := range urls {
		if u.RawQuery != "" {
			scrubbedURL := strings.Replace(u.String(), "?"+u.RawQuery, "", 1)
			finalText = strings.Replace(finalText, u.String(), scrubbedURL, 1)
		}
	}

	return finalText
}

func getURLsFromMessage(message *tgbotapi.Message) []*url.URL {
	if len(message.Entities) == 0 || message.Text == "" {
		return nil
	}

	var urlEntities []tgbotapi.MessageEntity
	for _, entity := range message.Entities {
		if entity.Type == "url" {
			urlEntities = append(urlEntities, entity)
		}
	}

	if len(urlEntities) == 0 {
		return nil
	}

	var urls []*url.URL
	for _, entity := range urlEntities {
		urlStr := message.Text[entity.Offset : entity.Offset+entity.Length]
		if u, err := url.Parse(urlStr); err == nil {
			urls = append(urls, u)
		}
	}

	return urls
}

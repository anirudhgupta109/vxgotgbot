package helpers

import (
	"fmt"
	"net/url"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StripURLs(message *tgbotapi.Message) string {
	if message.Text == "" {
		return ""
	}

	urls := getURLsFromMessage(message)
	finalText := ""

	for idx, u := range urls {
		finalText = fmt.Sprintf("%s %s", finalText, u)
		if u.RawQuery != "" {
			scrubbedURL := strings.Replace(u.String(), "?"+u.RawQuery, "", 1)
			finalText = fmt.Sprintf("%s\n[%d] %s", finalText, idx, scrubbedURL)
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

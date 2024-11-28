package socials

import (
	"regexp"
	"tgbot/helpers"
)

func HandleTikTok(message string) string {

	var matchRegex = regexp.MustCompile(helpers.TikTokRegex)

	matches := matchRegex.FindStringSubmatch(message)
	if matches == nil {
		return message
	}

	hostIndex := matchRegex.SubexpIndex(hostMatchGroup)
	message = regexp.MustCompile(matches[hostIndex]).ReplaceAllString(message, "vxtiktok.com")

	return message
}

package socials

import (
	"regexp"
	"tgbot/helpers"
)

func HandleInstagram(message string) string {

	var matchRegex = regexp.MustCompile(helpers.InstaRegex)

	matches := matchRegex.FindStringSubmatch(message)
	if matches == nil {
		return message
	}

	hostIndex := matchRegex.SubexpIndex(hostMatchGroup)
	message = regexp.MustCompile(matches[hostIndex]).ReplaceAllString(message, "ddinstagram.com")

	return message
}

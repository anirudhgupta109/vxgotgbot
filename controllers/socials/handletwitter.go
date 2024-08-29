package socials

import (
	"regexp"
	"strings"
	"tgbot/helpers"
)

const (
	hostMatchGroup = "host"
	rootMatchGroup = "root"
)

func HandleTwitter(message string) string {

	var matchRegex = regexp.MustCompile(helpers.TwitterRegex)

	matches := matchRegex.FindStringSubmatch(message)
	if matches == nil {
		return message
	}

	hostIndex := matchRegex.SubexpIndex(hostMatchGroup)
	rootIndex := matchRegex.SubexpIndex(rootMatchGroup)

	switch matches[rootIndex] {
	case "twitter":
		message = strings.Replace(message, matches[hostIndex], "vxtwitter.com", 1)

	case "x":
		message = strings.Replace(message, matches[hostIndex], "fixupx.com", 1)
	}

	return message
}

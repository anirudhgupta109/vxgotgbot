package helpers

const InstaRegex string = `https://(?:www\\.)?(?P<host>instagram\\.com)/(.*/)?(p|reel|tv)/[A-Za-z0-9]+.*`

const TwitterRegex string = `https://(?P<host>(?:mobile\.)?(?P<root>(twitter|x))\.com)/.*/status/[0-9]+.*`

func GetDomains() []string {
	return []string{"https://www.instagram.com/", "https://twitter.com/", "https://x.com/"}
}

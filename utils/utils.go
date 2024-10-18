package utils

func ParseURL(url string) string {
	if url[:len("https:/")] == "https:/" {
		url = "https://" + url[len("https:/"):]
	}
	return url
}

func IsWikiURL(url string) bool {
	return url[:len("https://en.wikipedia.org/wiki/")] == "https://en.wikipedia.org/wiki/"
}
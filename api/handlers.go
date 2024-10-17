package api

import (
	"net/http"

	"github.com/alexleyoung/summarizer/scraper"
)

func Scrape(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/"):]
	if url[:len("https:/")] == "https:/" {
		url = "https://" + url[len("https:/"):]
	}
	page := scraper.ScrapeURL(url)
	w.Write([]byte(page.Titles[0]))
	for _, p := range page.Paragraphs {
		w.Write([]byte(p))
	}
}

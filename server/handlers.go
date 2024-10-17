package server

import (
	"net/http"
	"os"

	"github.com/alexleyoung/summarizer/logic"
)

func Scrape(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/"):]
	if url[:len("https:/")] == "https:/" {
		url = "https://" + url[len("https:/"):]
	}
	
	page := logic.ScrapeGeneric(url)
	content := page.Titles[0] + "\n\n"
	for _, paragraph := range page.Paragraphs {
		content += paragraph + "\n\n"
	}

	respone := logic.Chat(os.Getenv("OPENAI_API_KEY"), content)
	w.Write([]byte(respone))
}

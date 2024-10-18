package server

import (
	"errors"
	"fmt"
	"io"
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

	stream, err := logic.ChatStream(os.Getenv("OPENAI_API_KEY"), content); if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer stream.Close()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.(http.Flusher).Flush()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(response.Choices[0].Delta.Content))
		w.(http.Flusher).Flush()
	}
}

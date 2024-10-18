package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alexleyoung/summarizer/logic"
	"github.com/alexleyoung/summarizer/utils"
)

func Scrape(w http.ResponseWriter, r *http.Request) {
	url := utils.ParseURL(r.URL.Path[len("/"):])

	content := ""
	
	if utils.IsWikiURL(url) {
		content = logic.ScrapeWiki(url).Content
	} else {
		content = logic.ScrapeGeneric(url).Content
	}


	streamOutput(w, content)
}

func streamOutput(w http.ResponseWriter, content string) {
	stream, err := logic.ChatStream(os.Getenv("OPENAI_API_KEY"), content); if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer stream.Close()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.(http.Flusher).Flush()

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

func returnStream(w http.ResponseWriter, content string) {
	stream, err := logic.ChatStream(os.Getenv("OPENAI_API_KEY"), content); if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer stream.Close()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.(http.Flusher).Flush()

	// Flush the headers to ensure they are sent immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return // End of stream
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Format the output for SSE: data: <content>\n\n
		msg := fmt.Sprintf("data: %s\n\n", response.Choices[0].Delta.Content)
		w.Write([]byte(msg))

		// Flush to ensure the client gets the message immediately
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}
}
package api

import (
	"fmt"
	"net/http"
)

func Scrape(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/"):]
	fmt.Fprintf(w, "Scraping %s", url)
}

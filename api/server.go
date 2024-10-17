package api

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", Scrape)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
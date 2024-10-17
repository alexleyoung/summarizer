package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexleyoung/summarizer/api"
)


func main() {
	
	http.HandleFunc("/", api.Scrape)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
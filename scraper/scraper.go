package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Init() {
	c := colly.NewCollector()
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong", err)
	})
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Found title:", e.Text)
	})
}
package logic

import (
	"fmt"

	"github.com/alexleyoung/summarizer/utils"
	"github.com/gocolly/colly"
)

func ScrapeURL(url string) utils.Page {
	c := colly.NewCollector()

	page := utils.Page{}
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong", err)
	})
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		page.Titles = append(page.Titles, e.Text)
	})
	c.OnHTML("p", func(e *colly.HTMLElement) {
		page.Paragraphs = append(page.Paragraphs, e.Text)
	})

	c.Visit(url)
	return page
}
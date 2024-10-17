package logic

import (
	"fmt"

	"github.com/alexleyoung/summarizer/utils"
	"github.com/gocolly/colly"
)

func ScrapeGeneric(url string) utils.GenericPage {
	c := colly.NewCollector()

	page := utils.GenericPage{}
	
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

func ScrapeWiki(url string) utils.WikiPage {
	c := colly.NewCollector()

	page := utils.WikiPage{}
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong", err)
	})
	c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
		page.Title = e.Text
	})
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		page.Headings = append(page.Headings, e.Text)
	})
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		page.Subtitles = append(page.Subtitles, e.Text)
	})
	c.OnHTML("p", func(e *colly.HTMLElement) {
		page.Paragraphs = append(page.Paragraphs, e.Text)
	})

	c.Visit(url)
	return page
}
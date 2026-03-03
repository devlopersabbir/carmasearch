package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func Scrape(url string) (interface{}, error) {
	if url == "" {
		return nil, nil
	}

	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
	)

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		fmt.Println("Title:::::::::::", e.Text)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Fail to scrape data")
	})

	err := c.Visit(url)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

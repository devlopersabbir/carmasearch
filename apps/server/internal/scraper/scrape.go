package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func Scrape(url string) error {
	if url == "" {
		return nil
	}

	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
	)

	c.OnHTML("a", HandleScrape)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		return err
	}

	return nil
}

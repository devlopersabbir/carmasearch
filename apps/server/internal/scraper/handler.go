package scraper

import "github.com/gocolly/colly/v2"

func HandleScrape(e *colly.HTMLElement) {
	e.Request.Visit(e.Attr("href"))
}

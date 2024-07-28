package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnHTML("div.ipc-metadata-list-summary-item__tc", func(e *colly.HTMLElement) {
		title := e.ChildText("a.ipc-title-link-wrapper h3.ipc-title__text")
		spans := e.DOM.Find("span.sc-b189961a-8.hCbzGp.cli-title-metadata-item")
		year := spans.Eq(0).Text()
		genre := spans.Eq(1).Text()
		rating := spans.Eq(2).Text()
		score := e.ChildText("div.sc-e2dbc1a3-0 jeHPdh sc-b189961a-2 bglYHz cli-ratings-container span.ipc-rating-star--rating")
		rawText := e.ChildText("div.sc-e2dbc1a3-0 jeHPdh sc-b189961a-2 bglYHz cli-ratings-container span.ipc-rating-star--voteCount")

		cleanText := strings.TrimSpace(strings.Trim(rawText, "(&nbsp;())"))
		fmt.Printf("Movie Title: %s, Year: %s, Duration: %s, Rating: %s, Score: %s, Reviews: %s\n", title, year, genre, rating, score, cleanText)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Connected to URL: ", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scrapping is done on: ", r.Request.URL)
	})
	c.Visit("https://www.imdb.com/chart/top/")
}

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnHTML("div.ipc-metadata-list-summary-item__tc div.sc-b189961a-0.iqHBGn.cli-children", func(e *colly.HTMLElement) {
		title := e.ChildText("a.ipc-title-link-wrapper h3.ipc-title__text")
		split := strings.SplitN(title, " ", 2)
		cleanTitle := strings.TrimSpace(split[1])
		spans := e.DOM.Find("span.sc-b189961a-8.hCbzGp.cli-title-metadata-item")
		year := spans.Eq(0).Text()
		genre := spans.Eq(1).Text()
		rating := spans.Eq(2).Text()
		score := e.ChildText("span.ipc-rating-star--rating")
		rawText := e.ChildText("span.ipc-rating-star--voteCount")

		cleanText := strings.TrimSpace(strings.Trim(rawText, "(&nbsp;())"))
		fmt.Printf("\nMovie Title: %s, Year: %s, Duration: %s, Rating: %s, Score: %s, Reviews: %s \n", cleanTitle, year, genre, rating, score, cleanText)

	})
	c.OnHTML("a.lister-page-next.next-page", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		e.Request.Visit(e.Request.AbsoluteURL(nextPage))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Connected to URL: ", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scrapping is done on: ", r.Request.URL, time.Now().Local().Format(time.RFC822))
	})
	c.Visit("https://www.imdb.com/chart/top/")
}

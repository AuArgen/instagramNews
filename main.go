package main

import (
	// "fmt"

	"github.com/gocolly/colly"
)

func main() {
	scrapPage("http://knews.kg")
}

func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("div", func(e *colly.HTMLElement) {
		enWord := e.DOM.Find(".td-block-row > .td-block-span4").Text()
		println("enWord = ",enWord)
	})

	c.Visit(url)
	// c.Visit("https://www.en365.ru/")
	// c.Visit("https://avn.oshsu.kg:85/")
}
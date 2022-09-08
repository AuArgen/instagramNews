package main

import (
	// "fmt"
	"time"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	scrapPage("http://knews.kg")
}

func scrapPage(url string) {
	c := colly.NewCollector()
	var count int;
	count = 0;
	var counton int = 0
	c.OnHTML("div a", func(e *colly.HTMLElement) {
		var link string = e.Attr("href")
		// if enWord != "" && enWord[0] != '<'{
		// 	println(count, " ) ", enWord) 
		// } else if enWord != "" {
		// 	count += 1
		// 	var a int = 0
		// 	var b int = 0
		// 	var s string = ""
		// 	for i := 0; i < len(enWord); i++ {
		// 		if enWord[i] == '<' {
		// 			a = 1
		// 		}
		// 		if a == 1 && b == 1 {
		// 			s = enWord[i:len(enWord)]
		// 			break
		// 		}
		// 		if enWord[i] =='>' {
		// 			b = 1
		// 		}
		// 	}
		// 	println(count , " ) ",s)
		// }
		curentTime := time.Now()
		curent := curentTime.Format("2006/01/02")
		var ss string = "https://knews.kg/"+curent
		if strings.Count(link,ss) >= 1 && counton%2 == 0 {
			count += 1
			scrapPageInPage(link,count)
		}
		counton += 1
	})

	c.Visit(url)
	// c.Visit("https://www.en365.ru/")
	// c.Visit("https://avn.oshsu.kg:85/")
}
func scrapPageInPage(url string, count int) {
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		var h1 string = e.DOM.Find("header>h1").Text()
		var p string = e.DOM.Find(".td-post-content p").Text() 
		println(count , " ) " ,url, " - ", h1)
		println(p)
		println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		println(" ")
	})

	c.Visit(url)
}
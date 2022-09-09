package main

import(
	"fmt"
	"net/http"
	"time"
	"strings"

	"github.com/gocolly/colly"
	
)
var ans = "</div> <h1>Seaerch information, pleace wait or update page!</h1><hr>"
func scrapPage(url string) {
		c := colly.NewCollector()
		var count int;
		count = 0;
		var counton int = 0
		c.OnHTML("div a", func(e *colly.HTMLElement) {
			var link string = e.Attr("href")
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
	}
	func scrapPageInPage(url string, count int) {
		c := colly.NewCollector()
		c.OnHTML("body", func(e *colly.HTMLElement) {
			var h1 string = e.DOM.Find("header>h1").Text()
			var p string = e.DOM.Find(".td-post-content p").Text() 
			ans += fmt.Sprintf("<p>%d) <a href='%s'>%s</a></p>", count,url,h1)
			ans += "<p>" +p +"</p> <hr>"
		})

		c.Visit(url)
	}

func home_page(w http.ResponseWriter, r * http.Request) {
	scrapPage("http://knews.kg")
	fmt.Fprintf(w,"<h1> Go is working!</h1> <hr/> <a href=/contacts/>Contacts</a> <hr><div style='display:none;'>",ans)
}

func contacts_page(w http.ResponseWriter, r*http.Request) {
	fmt.Fprintf(w,"<h1>Contacts paga working</h1> <hr>")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8083", nil)
}

func main() {
	handleRequest()
}
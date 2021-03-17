package sample_1

import (
	"fmt"
	"github.com/gocolly/colly"
)

const (
	DEMO_SITE = "https://webscraper.io/test-sites/e-commerce/allinone"
)

func Run() {
	collector := colly.NewCollector(
		colly.MaxDepth(1),
	)

	results := []string{}

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL.String())
	})

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		class := e.Attr("class")
		if class == "title" {
			link := e.Attr("href")
			e.Request.Visit(link)
		}
	})

	collector.OnHTML("div[class=caption]", func(e *colly.HTMLElement) {
		tempText := make([]string, 2)
		e.ForEach("h4", func(_ int, element *colly.HTMLElement) {
			tempText = append(tempText, element.Text)
		})
		fmt.Println(tempText)
		results = append(results, tempText[0])
		results = append(results, tempText[1])
	})

	collector.Visit("https://webscraper.io/test-sites/e-commerce/allinone/product/582")

	for _, url := range results {
		fmt.Println("\t", url)
	}
	fmt.Println(len(results))
}

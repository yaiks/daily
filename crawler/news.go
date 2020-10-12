package crawler

import (
	"github.com/gocolly/colly/v2"
)

// Information type
type Information struct {
	Title  string
	Href   string
	Domain string
}

// HandleBrazilNews is a method from Crawler type
func (crawler Crawler) HandleBrazilNews() (info []Information) {
	link := "https://www.globo.com"

	crawler.C.OnHTML(".hui-premium.hui-color-journalism", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		element := goquerySelection.Find(" a")
		href, _ := element.Attr("href")
		title := element.Children().Text()

		info = append(info, Information{
			Title:  title,
			Href:   href,
			Domain: "globo",
		})
	})

	crawler.C.Visit(link)

	return info
}

// "bra": "https://www.globo.com",
// "usa": "https://www.nytimes.com/section/us",

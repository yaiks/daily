package news

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
func HandleBrazilNews() (info []Information) {
	globoCollector := colly.NewCollector()
	// uolCollector := colly.NewCollector()

	globoCollector.OnHTML(".hui-premium.hui-color-journalism", func(e *colly.HTMLElement) {
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

	globoCollector.Visit("https://www.globo.com")

	return info
}

// "bra": "https://www.globo.com",
// "usa": "https://www.nytimes.com/section/us",

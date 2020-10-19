package news

import (
	"fmt"

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
	uolCollector := colly.NewCollector()

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

	uolCollector.OnHTML("a.manchete-editorial", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		title := e.ChildText("h1")

		info = append(info, Information{
			Title:  title,
			Href:   href,
			Domain: "uol",
		})
	})

	uolCollector.OnHTML("div.submanchete-destaque", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")
		title := e.ChildText("h2")

		info = append(info, Information{
			Title:  title,
			Href:   href,
			Domain: "uol",
		})
	})

	globoCollector.Visit("https://www.globo.com")
	uolCollector.Visit("https://www.uol.com.br/")

	return info
}

// HandleUSANews is a method from Crawler type
func HandleUSANews() (info []Information) {
	nytCollector := colly.NewCollector()

	nytCollector.OnHTML("section[id=collection-highlights-container]", func(e *colly.HTMLElement) {
		e.ForEach("h2 > a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			title := elem.Text

			info = append(info, Information{
				Title:  title,
				Href:   fmt.Sprintf("https://www.nytimes.com%s", href),
				Domain: "new york times",
			})
		})
	})

	nytCollector.Visit("https://www.nytimes.com/section/us")

	return info
}

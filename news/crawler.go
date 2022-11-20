package news

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
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

	globoCollector.OnHTML(".area-destaque", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		element := goquerySelection.Find(".post__link")

		element.Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			title := s.Children().Text()

			info = append(info, Information{
				Title:  title,
				Href:   href,
				Domain: "globo",
			})
		})
	})

	globoCollector.Visit("https://www.globo.com")

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

package countries

import (
	"github.com/gocolly/colly/v2"
)

// Country type
type Country struct {
	C *colly.Collector
}

// News type
type News struct {
	Title  string
	Href   string
	Domain string
}

// HandleBrazil is a method from Country type
func (country Country) HandleBrazil() (news []News) {
	link := "https://www.globo.com"

	country.C.OnHTML(".hui-premium.hui-color-journalism", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		element := goquerySelection.Find(" a")
		href, _ := element.Attr("href")
		title := element.Children().Text()

		news = append(news, News{
			Title:  title,
			Href:   href,
			Domain: "globo",
		})
	})

	country.C.Visit(link)

	return news
}

// MapCommandCountry maps a country to a function
func MapCommandCountry() map[string]func(*Country) (news []News) {
	m := map[string]func(*Country) (news []News){
		"bra": (*Country).HandleBrazil,
	}

	return m
}

// "bra": "https://www.globo.com",
// "usa": "https://www.nytimes.com/section/us",

package countries

import (
	"github.com/gocolly/colly/v2"
)

// Country type
type Country struct {
	C *colly.Collector
}

// News type
// type News struct {
// 	title string
// 	href  string
// }

// HandleBrazil is a method from Country type
func (country Country) HandleBrazil() (title []string, href []string) {
	link := "https://www.globo.com"

	country.C.OnHTML(".hui-premium.hui-color-journalism", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		element := goquerySelection.Find(" a")
		crawlerHref, _ := element.Attr("href")
		crawlerTitle := element.Children().Text()

		title = append(title, crawlerTitle)
		href = append(href, crawlerHref)
	})

	country.C.Visit(link)

	return title, href
}

// MapCommandCountry maps a country to a function
func MapCommandCountry() map[string]func(*Country) (title []string, href []string) {
	m := map[string]func(*Country) (title []string, href []string){
		"bra": (*Country).HandleBrazil,
	}

	return m
}

// "bra": "https://www.globo.com",
// "usa": "https://www.nytimes.com/section/us",

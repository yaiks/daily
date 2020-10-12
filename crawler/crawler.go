package crawler

import "github.com/gocolly/colly/v2"

// Crawler type
type Crawler struct {
	C *colly.Collector
}

// MapCommandCountry maps a country to a function
func MapCommandCountry(subject string) map[string]func(*Crawler) (info []Information) {
	if subject == "news" {
		m := map[string]func(*Crawler) (info []Information){
			"bra": (*Crawler).HandleBrazilNews,
		}

		return m
	}

	return nil
}

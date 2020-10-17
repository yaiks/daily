package news

// MapCommandCountry maps a country to a function
func MapCommandCountry(subject string) map[string]func() (info []Information) {
	if subject == "news" {
		m := map[string]func() (info []Information){
			"bra": HandleBrazilNews,
		}

		return m
	}

	return nil
}

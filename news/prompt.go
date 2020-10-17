package news

import (
	"strings"

	"github.com/manifoldco/promptui"
)

// HandleNewsPrompt handles the format and display of news in CLI
func HandleNewsPrompt(info []Information) (int, string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Title }}?",
		Active:   "\U0001f449 {{ .Title | cyan }}",
		Inactive: "  {{ .Title | white }}",
		Selected: "\U0001f449 {{ .Title | cyan }}",
		Details: `
		--------- News ----------
		{{ "Title:" | faint }}	{{ .Title }}
		{{ "Domain:" | faint }}	{{ .Domain }}`,
	}

	searcher := func(input string, index int) bool {
		news := info[index]
		title := strings.Replace(strings.ToLower(news.Title), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(title, input)
	}

	prompt := promptui.Select{
		Label:     "See the news for today",
		Items:     info,
		HideHelp:  true,
		Templates: templates,
		Searcher:  searcher,
	}

	return prompt.Run()
}

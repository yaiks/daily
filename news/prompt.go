package news

import (
	"github.com/manifoldco/promptui"
)

// HandleNewsPrompt handles the format and display of news in CLI
func HandleNewsPrompt(info []Information) (int, string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Title }}?",
		Active:   "\U0001f449 {{ .Title | cyan }}",
		Inactive: "  {{ .Title | white }}",
		Selected: "\U0001f449 {{ .Title | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Check the news for today",
		Items:     info,
		HideHelp:  true,
		Templates: templates,
	}

	return prompt.Run()
}

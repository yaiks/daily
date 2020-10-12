package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/manifoldco/promptui"
	"github.com/ricardohan93/daily/crawler"
	"github.com/ricardohan93/daily/utils"
	"github.com/spf13/cobra"
)

func customValidation(cms *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("require a country. Type **daily list countries** to see the list of available values")
	}

	if len(args) > 1 {
		return errors.New("command *news* accept only one country. Type **daily list countries** to see the list of available values")
	}

	// if _, ok := countries.Countries[args[0]]; !ok {
	// 	return errors.New("we don't support this country yet. Type **daily list countries** to see the list of available values. You are welcome to open a PR to add this country to our list")
	// }

	return nil
}

func newsCommand(cmd *cobra.Command, args []string) {
	countryInput := args[0]

	var info []crawler.Information

	c := colly.NewCollector()
	cwlr := crawler.Crawler{C: c}

	mapCrawlerFunction := crawler.MapCommandCountry("news")

	if fn := mapCrawlerFunction[countryInput]; fn != nil {
		info = fn(&cwlr)
	}

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

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	utils.OpenBrowser(info[index].Href)
}

// newsCmd represents the news command
var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "displays news according to the country you choose",
	Long:  `news is a daily subcommand to display relevant news for you`,
	Args:  customValidation,
	Run:   newsCommand,
}

func init() {
	rootCmd.AddCommand(newsCmd)
}

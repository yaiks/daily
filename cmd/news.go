package cmd

import (
	"errors"
	"fmt"

	"github.com/ricardohan93/daily/news"
	"github.com/ricardohan93/daily/utils"
	"github.com/spf13/cobra"
)

var countries = map[string]string{
	"bra": "Brazil",
}

func customValidation(cms *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("require a country. Type **daily list countries** to see the list of available values")
	}

	if len(args) > 1 {
		return errors.New("command *news* accept only one country. Type **daily list countries** to see the list of available values")
	}

	if _, ok := countries[args[0]]; !ok {
		return errors.New("we don't support this country yet. Type **daily list countries** to see the list of available values. You are welcome to open a PR to add this country to our list")
	}

	return nil
}

func newsCommand(cmd *cobra.Command, args []string) {
	countryInput := args[0]

	var info []news.Information

	mapCrawlerFunction := news.MapCommandCountry("news")

	if fn := mapCrawlerFunction[countryInput]; fn != nil {
		info = fn()
	}

	index, _, err := news.HandleNewsPrompt(info)

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

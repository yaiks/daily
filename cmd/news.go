package cmd

import (
	"fmt"

	"github.com/ricardohan93/daily/news"
	"github.com/ricardohan93/daily/utils"
	"github.com/spf13/cobra"
)

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
	Use:   "news <country>",
	Short: "displays news according to the country you choose",
	Long:  `news is a daily subcommand to display relevant news for you`,
	Args:  cobra.ExactArgs(1),
	Run:   newsCommand,
}

func init() {
	rootCmd.AddCommand(newsCmd)
}

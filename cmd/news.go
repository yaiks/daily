package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/manifoldco/promptui"
	"github.com/ricardohan93/daily/constants"
	"github.com/spf13/cobra"
)

func customValidation(cms *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("require a country. Type **daily list countries** to see the list of available values")
	}

	if len(args) > 1 {
		return errors.New("command *news* accept only one country. Type **daily list countries** to see the list of available values")
	}

	if _, ok := constants.Countries[args[0]]; !ok {
		return errors.New("we don't support this country yet. Type **daily list countries** to see the list of available values. You are welcome to open a PR to add this country to our list")
	}

	return nil
}

var urls = []string{"https://www.globo.com", "https://www.folha.com.br"}

func createFileName(fullURL string) (host string) {
	regx := regexp.MustCompile(`(www\.|\.com|\.br)`)
	u, err := url.Parse(fullURL)
	if err != nil {
		panic(err)
	}

	host = u.Hostname()
	host = regx.ReplaceAllString(host, "")

	return host
}

func newsCommand(cmd *cobra.Command, args []string) {
	// country := args[0]

	c := colly.NewCollector()

	var items []string

	c.OnHTML(".hui-premium.hui-color-journalism", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM

		element := goquerySelection.Find(" a")
		title := element.Children().Text()
		// link, _ := element.Attr("href")

		// fmt.Println("TITLE - ", title)
		// fmt.Println("LINK - ", link)

		items = append(items, title)
	})

	c.Visit("https://www.globo.com")

	prompt := promptui.Select{
		Label: "See the news for today",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

// newsCmd represents the news command
var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "Bringing relevant news for you",
	Long:  `news is a daily subcommand to display relevant news for you`,
	Args:  customValidation,
	Run:   newsCommand,
}

func init() {
	rootCmd.AddCommand(newsCmd)
}

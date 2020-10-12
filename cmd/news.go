package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/manifoldco/promptui"
	"github.com/ricardohan93/daily/countries"
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
	userCountry := args[0]

	var title []string
	var href []string

	c := colly.NewCollector()
	country := countries.Country{C: c}

	m := countries.MapCommandCountry()

	if fn := m[userCountry]; fn != nil {
		title, href = fn(&country)
	}

	prompt := promptui.Select{
		Label: "See the news for today",
		Items: title,
	}

	int, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	utils.OpenBrowser(href[int])
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

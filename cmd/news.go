package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newsCmd represents the news command
var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "Bringing relevant news for you",
	Long:  `news is a daily subcommand to display relevant news for you`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("news called")
	},
}

func init() {
	rootCmd.AddCommand(newsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "daily",
	Short: "Bringing relevant information for you",
	Long:  `Daily is a program to keep yourself informed without leaving your terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		countries, _ := cmd.Flags().GetBool("countries")

		if countries {
			fmt.Println("bra --> Brazil\nusa --> United States of America")
		} else {
			fmt.Println("Welcome to daily!\nCheck our supported commands:\n\nnews {{ country prefix }}    see the news for a specific country")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.daily.yaml)")
	rootCmd.Flags().BoolP("countries", "c", false, "List of available countries")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".daily" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".daily")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

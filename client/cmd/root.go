package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var BaseURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "",
	Long: `
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		fmt.Println("+++++++++++++++++++++++++++++++++++++")
		os.Exit(1)
	}
}

func init()  {
	cobra.OnInitialize(initConfig) // initialize cobra with config

	BaseURL = os.Getenv("BASE_URL")

	fmt.Println("+++++++++++++++++++++++++++++++++++++")
}

// initConfig reads in config file and ENV variables if set.
// no required in our case
func initConfig() {
}

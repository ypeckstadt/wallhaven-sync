
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wallhaven-sync",
	Short: "Wallhaven.cc pictures sync CLI app",
	Long: `CLI app that downloads all the pictures in a Wallhaven.cc collection`,
	Run: func(cmd *cobra.Command, args []string) { fmt.Println("Wallhaven sync") },
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
	rootCmd.PersistentFlags().StringP("api-key", "k", "", "Wallpaper.cc API key")
	rootCmd.MarkPersistentFlagRequired("api-key")
}
package cmd

import (
	"context"
	"fmt"

	"github.com/mpppk/go-scrapbox/scrapbox"
	"github.com/mpppk/sbox/utl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show page contents",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		defaultProjectName := viper.GetString("project")
		defaultServerName := viper.GetString("server")
		targetPage := utl.ParsePagePath(args[0], defaultProjectName, defaultServerName)
		client := scrapbox.NewClient(nil)
		contents, _, err := client.Pages.GetText(context.Background(), targetPage.Project, targetPage.Name)
		if err != nil {
			fmt.Println("failed to fetch page from " + targetPage.String())
		}

		fmt.Println(contents)
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}

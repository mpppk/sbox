package cmd

import (
	"context"
	"fmt"

	"github.com/mpppk/go-scrapbox/scrapbox"
	"github.com/mpppk/sbox/utl"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show page contents",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		targetPage := utl.ParsePagePath(args[0], projectName, serverName)
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

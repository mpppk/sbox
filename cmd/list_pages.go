package cmd

import (
	"context"
	"fmt"

	"github.com/mpppk/go-scrapbox/scrapbox"
	"github.com/mpppk/sbox/utl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var limit int

var listPagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "list pages",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultProjectName := viper.GetString("project")
		defaultServerName := viper.GetString("server")
		targetPage := utl.ParsePagePath("dummy", defaultProjectName, defaultServerName)
		client := scrapbox.NewClient(nil)
		pages, _, err := client.Pages.ListByProject(context.Background(), targetPage.Project,
			&scrapbox.PageListByProjectOptions{Limit: limit})

		if err != nil {
			fmt.Println("failed to fetch pages from " + targetPage.String())
		}

		for _, page := range pages {
			fmt.Println(page.Title)
		}
	},
}

func init() {
	listPagesCmd.PersistentFlags().IntVar(&limit, "limit", 100, "pages num limit")
	listCmd.AddCommand(listPagesCmd)
}

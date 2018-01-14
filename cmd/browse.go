package cmd

import (
	"fmt"
	"net/url"

	"github.com/mpppk/sbox/utl"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var contents string
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse page",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		page := utl.ParsePagePath(args[0], projectName, serverName)
		values := url.Values{}
		values.Add("body", contents)
		pageURLWithQuery := fmt.Sprintf("%s?%s", page.String(), values.Encode())
		open.Run(pageURLWithQuery)
	},
}

func init() {
	browseCmd.PersistentFlags().StringVar(&contents, "contents", "", "page contents")
	RootCmd.AddCommand(browseCmd)
}

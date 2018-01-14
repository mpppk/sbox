package cmd

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var contents string
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse page",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		names := strings.Split(args[0], "/")
		pageName := names[len(names)-1]
		if len(names) > 1 {
			projectName = names[len(names)-2]
		}
		if len(names) > 2 {
			serverName = names[len(names)-3]
		}

		values := url.Values{}
		values.Add("body", contents)
		pageURL := path.Join(serverName, projectName, pageName)
		pageURLWithQuery := fmt.Sprintf("%s?%s", pageURL, values.Encode())
		open.Run(pageURLWithQuery)
	},
}

func init() {
	browseCmd.PersistentFlags().StringVar(&contents, "contents", "", "page contents")
	RootCmd.AddCommand(browseCmd)
}

package cmd

import (
	"net/url"
	"strings"

	"github.com/mpppk/sbox/utl"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var contents string
var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse page",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		defaultProjectName := viper.GetString("project")
		defaultServerName := viper.GetString("server")
		tempTrimmedTitle := strings.Replace(args[0], "\r", "", -1)
		trimmedTitle := strings.Replace(tempTrimmedTitle, "\n", "", -1)
		escapedTitle := url.PathEscape(trimmedTitle)
		targetPage := utl.ParsePagePath(escapedTitle, defaultProjectName, defaultServerName)
		query := ""
		if contents != "" {
			values := url.Values{}
			values.Add("body", contents)
			query = "?" + values.Encode()
		}
		pageURLWithQuery := targetPage.String() + query
		open.Run(pageURLWithQuery)
	},
}

func init() {
	browseCmd.PersistentFlags().StringVar(&contents, "contents", "", "page contents")
	RootCmd.AddCommand(browseCmd)
}

package cmd

import (
	"fmt"
	"net/url"

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
		targetPage := utl.ParsePagePath(args[0], defaultProjectName, defaultServerName)
		values := url.Values{}
		values.Add("body", contents)
		pageURLWithQuery := fmt.Sprintf("%s?%s", targetPage.String(), values.Encode())
		open.Run(pageURLWithQuery)
	},
}

func init() {
	browseCmd.PersistentFlags().StringVar(&contents, "contents", "", "page contents")
	RootCmd.AddCommand(browseCmd)
}

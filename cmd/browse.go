package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "browse page",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("browse command is not impolemented yet")
	},
}

func init() {
	RootCmd.AddCommand(browseCmd)
}

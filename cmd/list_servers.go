package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listServersCmd represents the list_servers command
var listServersCmd = &cobra.Command{
	Use:   "servers",
	Short: "list servers",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list servers command is not impolemented yet")
	},
}

func init() {
	listCmd.AddCommand(listServersCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects, pages, and servers",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command is not impolemented yet")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

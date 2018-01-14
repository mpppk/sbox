package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listPagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "list pages",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list pages command is not impolemented yet")
	},
}

func init() {
	listCmd.AddCommand(listPagesCmd)
}

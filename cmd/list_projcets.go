package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "list projects",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list projects command is not impolemented yet")
	},
}

func init() {
	listCmd.AddCommand(listProjectsCmd)
}

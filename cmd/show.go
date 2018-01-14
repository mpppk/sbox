package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show page",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show command is not impolemented yet")
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}

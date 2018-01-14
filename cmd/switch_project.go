package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var switchProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "switch project",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch project command is not impolemented yet")
	},
}

func init() {
	switchCmd.AddCommand(switchProjectCmd)
}

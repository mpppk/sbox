package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var switchServerCmd = &cobra.Command{
	Use:   "switch_server",
	Short: "switch server",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch server command is not impolemented yet")
	},
}

func init() {
	switchCmd.AddCommand(switchServerCmd)
}

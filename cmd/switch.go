package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "switch project or server",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch command is not impolemented yet")
	},
}

func init() {
	RootCmd.AddCommand(switchCmd)
}

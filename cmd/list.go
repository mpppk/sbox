package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects, pages, and servers",
	Long:  ``,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

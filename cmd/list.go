package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the incomplete tasks",
	Long:  "Lists out all the tasks which are incomplete.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fake task list.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

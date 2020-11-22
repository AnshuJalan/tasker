package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new task",
	Long:  "Creates a new task and adds it to the underlying database.",
	Run: func(cmd *cobra.Command, args []string) {
		newTask := strings.Join(args, " ")
		fmt.Println("Added new task- ", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

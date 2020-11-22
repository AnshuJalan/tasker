package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do <task id>",
	Short: "Marks a task as complete",
	Long:  "Marks a task on the TODO list as complete.",
	RunE: func(cmd *cobra.Command, args []string) error {

		ids := make([]int, len(args))

		for i, arg := range args {
			task, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument- ", arg)
				return err
			}
			ids[i] = task
		}

		fmt.Println("Finished task(s)- ", ids)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

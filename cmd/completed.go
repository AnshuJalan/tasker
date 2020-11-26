package cmd

import (
	"fmt"
	"os"

	"github.com/anshujalan/tasker/db"
	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists completed tasks",
	Long:  "Lists all tasks completed within the last 24 hours",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListCompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have completed no task in the past 24 Hours. 😴")
			return
		}

		fmt.Printf("---TASKS---\n")
		for _, task := range tasks {
			y, m, d := task.Completed.Date()
			h := task.Completed.Hour()
			min := task.Completed.Minute()
			fmt.Printf("%s\t(%d/%d/%d  %d:%d)\n", task.Value, d, m, y, h, min)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}

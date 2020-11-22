package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anshujalan/tasker/db"
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

		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id > len(tasks) {
				fmt.Printf("Task with id %d does not exist.", id)
				continue
			}
			err := db.DeleteTask(tasks[id-1].Key)
			if err != nil {
				fmt.Println("Something went wrong: ", err)
				os.Exit(1)
			}
			fmt.Printf("Marked task %d as completed.\n", id)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

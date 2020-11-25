package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anshujalan/tasker/db"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a task from the list",
	Long:  "Removes the specified tasks from the todo list database.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ids := make([]int, len(args))
		for i, arg := range args {
			val, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Could not parse the argument %s", arg)
				return err
			}
			ids[i] = val
		}

		tasks, _ := db.ListTasks()

		for _, id := range ids {
			if id > len(tasks) {
				fmt.Printf("Task with id %d does not exist!", id)
				continue
			}
			err := db.RemoveTask(tasks[id-1].Key)
			if err != nil {
				fmt.Println("Something went wrong: ", err)
				os.Exit(1)
			}
			fmt.Printf("Removed task %d from the list.\n", id)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

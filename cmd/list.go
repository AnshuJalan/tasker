package cmd

import (
	"fmt"
	"os"

	"github.com/anshujalan/tasker/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the incomplete tasks",
	Long:  "Lists out all the tasks which are incomplete.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks, why not take a vacation? 🌞")
			return
		}

		fmt.Println("--Your Tasks--")
		fmt.Printf("ID\tTASK\n--\t----\n")
		for i, task := range tasks {
			fmt.Printf("%d\t%s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/anshujalan/tasker/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Creates a new task",
	Long:  "Creates a new task and adds it to the underlying database.",
	Run: func(cmd *cobra.Command, args []string) {
		newTask := strings.Join(args, " ")
		_, err := db.AddTask(newTask)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		fmt.Printf("Added new task '%s'.", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

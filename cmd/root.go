package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Tasker is a CLI task management application",
	Long:  "A simple & fast task management application, controllable through CLI.",
}

//Execute runs the command
func Execute() error {
	return rootCmd.Execute()
}

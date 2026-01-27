package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long:  "Add a task to the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Add(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Delete a given task from the todo list by its index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Delete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

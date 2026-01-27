package cmd

import "github.com/spf13/cobra"

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  "Marks the task of the given id as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Complete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

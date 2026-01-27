package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks of the todo list",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

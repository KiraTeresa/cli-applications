package cmd

import "github.com/spf13/cobra"

var listAll bool
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks of the todo list",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		List(listAll)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "List all tasks, no matter the status")
	rootCmd.AddCommand(listCmd)
}

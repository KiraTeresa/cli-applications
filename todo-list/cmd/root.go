package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todo-list",
	Short: "todo-list performs crud operations on a data file of tasks",
	Long:  "todo-list performs crud operations on a data file of tasks - add, list, complete, delete",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops - executing todo-list failed: %s\n", err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "math-operation",
	Short: "math-operation is a cli tool for performing basic mathematical operations",
	Long:  "math-operation is a cli tool for performing basic mathematical operations - addition, multiplication, division and subtraction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing math-operation '%s'\n", err)
		os.Exit(1)
	}
}

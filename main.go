package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A simple greet application",
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet someone",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Hello, %s!\n", name)
		return nil
	},
}

func init() {
	greetCmd.Flags().StringP("name", "n", "world", "Name of the person to greet")
	rootCmd.AddCommand(greetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

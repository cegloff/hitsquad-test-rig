package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "hellotool",
	Short:   "A simple hello tool",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hellotool %s\n", version)
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets the user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, %s!\n", name)
	},
}

var name string

func init() {
	rootCmd.SetVersionTemplate("hellotool {{.Version}}\n")
	greetCmd.Flags().StringVarP(&name, "name", "n", "world", "name to greet")
	rootCmd.AddCommand(greetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

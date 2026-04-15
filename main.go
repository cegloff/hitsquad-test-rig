package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"

func main() {
	var versionBool bool
	var rootCmd = &cobra.Command{
		Use:   "hellotool",
		Short: "A simple hello tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hellotool %s\n", version)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&versionBool, "version", "v", false, "print version")

	var name string
	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greets the user",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	greetCmd.Flags().StringVarP(&name, "name", "n", "world", "name to greet")

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "hellotool",
		Version: "0.1.0",
		Short: "hellotool - A CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hellotool v0.1.0")
			os.Exit(0)
		},
	}

	cmd.Execute()
}

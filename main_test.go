package main

import (
	"bytes"
	"testing"
)

func executeCommand(args []string) (string, error) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	
	// Important: Reset flags for each test execution
	// Since cobra commands are global in this implementation, 
	// we should reset the name flag to default
	greetCmd.Flags().Set("name", "world")
	
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return buf.String(), err
}

func TestGreetCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Default name",
			args:     []string{"greet"},
			expected: "Hello, world!\n",
		},
		{
			name:     "Custom name",
			args:     []string{"greet", "--name", "Alice"},
			expected: "Hello, Alice!\n",
		},
		{
			name:     "Short flag",
			args:     []string{"greet", "-n", "Bob"},
			expected: "Hello, Bob!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := executeCommand(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if output != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output)
			}
		})
	}
}

func TestRootCommand(t *testing.T) {
	// Testing that root command doesn't crash and shows help
	// We can't easily check the full help text but just that it executes
	_, err := executeCommand([]string{})
	if err != nil {
		t.Errorf("root command failed with error: %v", err)
	}
}

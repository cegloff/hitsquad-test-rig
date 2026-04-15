package main

import (
	"os/exec"
	"strings"
	"testing"
)

func runCommand(args []string) string {
	cmd := exec.Command("./hellotool", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "Error: " + err.Error()
	}
	return strings.TrimSpace(string(out))
}

func TestHellotool(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Root command",
			args:     []string{},
			expected: "hellotool v0.1.0",
		},
		{
			name:     "Version flag",
			args:     []string{"--version"},
			expected: "hellotool v0.1.0",
		},
		{
			name:     "Greet default",
			args:     []string{"greet"},
			expected: "Hello, world!",
		},
		{
			name:     "Greet named",
			args:     []string{"greet", "--name", "Bob"},
			expected: "Hello, Bob!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runCommand(tt.args)
			if got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestGreetCommand(t *testing.T) {
	t.Run("default name", func(t *testing.T) {
		got := runCommand([]string{"greet"})
		if got != "Hello, world!" {
			t.Errorf("Expected %q, got %q", "Hello, world!", got)
		}
	})

	t.Run("custom name", func(t *testing.T) {
		got := runCommand([]string{"greet", "--name", "Bob"})
		if got != "Hello, Bob!" {
			t.Errorf("Expected %q, got %q", "Hello, Bob!", got)
		}
	})
}

package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestRootCommand tests the root command execution
func TestRootCommand(t *testing.T) {
	// Save original stdout
	originalOut := os.Stdout

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Redirect the command's output to our writer
	rootCmd.SetOut(w)
	rootCmd.SetErr(w)

	// Execute the root command with no arguments
	rootCmd.SetArgs([]string{})
	result := rootCmd.Execute()

	// Restore stdout
	w.Close()
	os.Stdout = originalOut

	// Check the command executed successfully
	if result != nil {
		t.Errorf("Expected root command to succeed, got error: %v", result)
	}

	// Read the captured output
	outBuf := &bytes.Buffer{}
	outBuf.ReadFrom(r)
	output := strings.TrimSpace(outBuf.String())
	expected := "hellotool v0.1.0"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestVersionFlag tests the version flag
func TestVersionFlag(t *testing.T) {
	// Save original stdout
	originalOut := os.Stdout

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Redirect the command's output
	rootCmd.SetOut(w)
	rootCmd.SetErr(w)

	// Execute with version flag
	rootCmd.SetArgs([]string{"--version"})
	result := rootCmd.Execute()

	// Restore stdout
	w.Close()
	os.Stdout = originalOut

	// Check success
	if result != nil {
		t.Errorf("Version flag execution failed: %v", result)
	}

	// Read the captured output
	outBuf := &bytes.Buffer{}
	outBuf.ReadFrom(r)
	output := strings.TrimSpace(outBuf.String())
	expected := "hellotool v0.1.0"
	if output != expected {
		t.Errorf("Expected version output %q, got %q", expected, output)
	}
}

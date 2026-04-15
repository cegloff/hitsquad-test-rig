package main

import (
	"os/exec"
	"testing"
)

func TestRootCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("Expected exit status 0, got error: %v", err)
	}

	expectedOutput := "hellotool v0.1.0\n"
	actualOutput := string(output)
	if actualOutput != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, actualOutput)
	}
}

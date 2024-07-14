// till_test.go
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func runCommand(args ...string) (string, int, error) {
	cmd := exec.Command("./till", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()
	return out.String(), exitCode, err
}

func TestTill(t *testing.T) {

	loc, err := time.LoadLocation("Europe/Prague")
	if err != nil {
		fmt.Println("Error loading location:", err)
		os.Exit(1)
	}

	// Get the current time plus one minute
	currentTime := time.Now().In(loc).Add(1 * time.Minute)
	timeString := currentTime.Format("15:04")

	tests := []struct {
		timeString   string
		shouldError  bool
		expectedCode int
	}{
		{timeString, false, 0}, // + 1 minute
		{"24:00", true, 1},     // invalid time
		{"invalid", true, 1},   // invalid format
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Time-%s", tt.timeString), func(t *testing.T) {
			output, exitCode, err := runCommand(tt.timeString)

			if exitCode != tt.expectedCode {
				t.Fatalf("expected exit code: %d, got: %d, output: %s", tt.expectedCode, exitCode, output)
			}

			if (err != nil) != tt.shouldError {
				t.Fatalf("expected error: %v, got: %v, output: %s", tt.shouldError, err, output)
			}

			if !tt.shouldError {
				expectedParts := []string{"Sleeping until", "that is in"}
				for _, part := range expectedParts {
					if !strings.Contains(output, part) {
						t.Errorf("expected output to contain %q, but it didn't. Full output: %s", part, output)
					}
				}
			} else {
				expectedParts := []string{"Error parsing time"}
				for _, part := range expectedParts {
					if !strings.Contains(output, part) {
						t.Errorf("expected output to contain %q, but it didn't. Full output: %s", part, output)
					}
				}
			}
		})
	}
}

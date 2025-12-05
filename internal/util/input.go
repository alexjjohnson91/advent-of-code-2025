package util

import (
	"fmt"
	"os"
)

// ReadFile reads the entire file at the given path and returns its contents
// as a string.
//
// It panics if the file cannot be read, wrapping the underlying error with
// context about the path. Use this helper in contexts where failing fast is
// acceptable (e.g., tests, small utilities, or startup-time config loading).
//
// Example:
//
//	input := ReadFile("data/input.txt")
//
func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read input %s: %w", path, err))
	}
	return string(b)
}

package util

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read input %s: %w", path, err))
	}
	return string(b)
}

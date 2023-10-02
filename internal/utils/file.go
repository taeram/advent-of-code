package utils

import (
	"os"
	"strings"
)

// ReadLines returns a string array of the given file contents.
func ReadLines(filePath string) []string {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(contents), "\n")
}

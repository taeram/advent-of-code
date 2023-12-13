package utils

import (
	"os"
	"strings"
)

// ReadInputFile returns an array of strings if the user passes in a filename to the command.
func ReadInputFile() []string {
	//nolint: gomnd
	if len(os.Args) < 2 {
		panic("No input file specified\n")
	}
	filePath := os.Args[1]

	contents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(contents), "\n"), "\n")
}

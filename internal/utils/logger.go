package utils

import "fmt"

// Logger logs to stdout.
func Logger(format string, a ...any) {
	//nolint: forbidigo
	fmt.Printf(format, a...)
	//nolint: forbidigo
	fmt.Printf("\n")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		showUsage()
		os.Exit(1)
	}

	if err := run(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "~wordcounter~\n")
	fmt.Fprintf(os.Stderr, "Usage: %s <path to file>\n", os.Args[0])
}

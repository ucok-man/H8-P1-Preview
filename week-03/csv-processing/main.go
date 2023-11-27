package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args[1:]) != 2 {
		showUsage()
		os.Exit(1)
	}

	if filepath.Ext(os.Args[1]) != ".csv" && filepath.Ext(os.Args[2]) != ".csv" {
		showUsage()
		os.Exit(1)
	}

	if err := run(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "~csv_processor~\n")
	fmt.Fprintf(os.Stderr, "Usage: %s <input file.csv> <output file.csv>\n", os.Args[0])
}

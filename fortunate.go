package main

import (
	"fmt"
	"os"

	"github.com/zquestz/fortunate/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.EntryCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

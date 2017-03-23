package main

import (
	_ "github.com/nkumar15/grep-feeds/matchers"
	"github.com/nkumar15/grep-feeds/search"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}

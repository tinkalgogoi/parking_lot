package main

import (
	"github.com/parking_lot/reader"
	"os"
)

func main() {
	args := os.Args[1:]

	// file commands
	if len(args) > 0 {
		fileName := args[0]
		reader.FileReader(fileName)
	} else {
		// stdin commands
		reader.StdInReader()
	}
}


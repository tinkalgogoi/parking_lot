package reader

import (
	"bufio"
	"fmt"
	"github.com/parking_lot/park"
	"io"
	"os"
	"strings"
)
// generic reader for types who implements io.Reader
func GenericReader(input io.Reader) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		park.ParkCar(command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Reading commands from stdin
func StdInReader() {
	GenericReader(os.Stdin)
}

// reading commands line by line from file
func FileReader(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	GenericReader(f)
}

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/afero"
	flag "github.com/spf13/pflag"
)

var inputFilename string

func init() {
	flag.StringVar(&inputFilename, "input", "", "Input filename")
}

func star1(lines []string) int {
	var sum int = 0
	for lineNumber, eachLine := range lines {
		eachLine = strings.TrimSpace(eachLine)
		if len(eachLine) == 0 {
			continue
		}
		// eachLine is a rucksack
		item := findMisplaced(eachLine)
		if item == SEARCH_ERROR {
			panic(fmt.Sprintf("search error on line %d: %s", lineNumber+1, eachLine))
		}
		sum += getValueOfItem(item)
	}
	return sum
}

func star2(lines []string) int {
	var sum int = 0
	for lineNumber := 0; lineNumber < len(lines)-2; lineNumber += 3 {
		var item rune = findBadge(
			strings.TrimSpace(lines[lineNumber]),
			strings.TrimSpace(lines[lineNumber+1]),
			strings.TrimSpace(lines[lineNumber+2]),
		)
		if item == SEARCH_ERROR {
			panic(fmt.Sprintf("search error on group over lines %d to %d", lineNumber+1, lineNumber+3))
		}
		sum += getValueOfItem(item)
	}
	return sum
}

func main() {
	var (
		err     error
		inputFs afero.Fs = afero.NewReadOnlyFs(afero.NewOsFs())
		rawData []byte
		lines   []string
	)

	flag.Parse()
	log.Printf("input filename: %#v\n", inputFilename)

	rawData, err = afero.ReadFile(inputFs, inputFilename)
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(rawData), "\n")
	fmt.Printf("Sum of priorities: %d\n", star1(lines))
	fmt.Printf("Sum of priorities: %d\n", star2(lines))
}

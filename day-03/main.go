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

func main() {
	var (
		err             error
		inputFs         afero.Fs = afero.NewReadOnlyFs(afero.NewOsFs())
		rawData         []byte
		lines           []string
		sumOfPriorities int
	)

	flag.Parse()
	log.Printf("input filename: %#v\n", inputFilename)

	rawData, err = afero.ReadFile(inputFs, inputFilename)
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(rawData), "\n")
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
		sumOfPriorities += getValueOfItem(item)
	}

	fmt.Printf("Sum of priorities: %d\n", sumOfPriorities)
}

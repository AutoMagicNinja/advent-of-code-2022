package main

import (
	"log"
	"strings"

	"github.com/spf13/afero"
	flag "github.com/spf13/pflag"
)

var inputFilename string

func init() {
	flag.StringVar(&inputFilename, "input", "", "Input filename")
}

func firstStar(lines []string) {
	var answer int = 0
	for lineNumber, eachLine := range lines {
		pair, err := NewPair(eachLine)
		if err != nil {
			log.Printf("error on line %d: %s", lineNumber+1, err.Error())
			continue
		}
		a, b := pair[0], pair[1]
		if a.Contains(b) || b.Contains(a) {
			// log.Printf("line contains: %s", eachLine)
			answer += 1
		}
	}
	log.Printf("first star: %d", answer)
}

func secondStar(lines []string) {
	var answer int = 0
	for lineNumber, eachLine := range lines {
		pair, err := NewPair(eachLine)
		if err != nil {
			log.Printf("error on line %d: %s", lineNumber+1, err.Error())
			continue
		}
		a, b := pair[0], pair[1]
		if a.Overlaps(b) {
			// log.Printf("line overlaps: %s", eachLine)
			answer += 1
		}
	}
	log.Printf("second star: %d", answer)
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

	lines = strings.Split(strings.TrimSpace(string(rawData)), "\n")
	firstStar(lines)
	secondStar(lines)
}

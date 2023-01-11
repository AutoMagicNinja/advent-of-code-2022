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
	var answer int
	for lineNumber, eachLine := range lines {
	}
	log.Printf("first star: %d", answer)
}

// func secondStar(lines []string) {
// 	var answer int
// 	for lineNumber, eachLine := range lines {
// 	}
// 	log.Printf("second star: %d", answer)
// }

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
	// secondStar(lines)
}

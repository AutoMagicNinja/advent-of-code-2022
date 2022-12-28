package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/automagicninja/advent-of-code-2022/day-02/rps"
	"github.com/spf13/afero"
	flag "github.com/spf13/pflag"
)

var inputFilename string

func init() {
	flag.StringVar(&inputFilename, "input", "", "Input filename")
}

func main() {
	var (
		err        error
		inputFs    afero.Fs = afero.NewReadOnlyFs(afero.NewOsFs())
		rawData    []byte
		lines      []string
		totalScore int
	)

	flag.Parse()
	log.Printf("input filename: %#v\n", inputFilename)

	rawData, err = afero.ReadFile(inputFs, inputFilename)
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(rawData), "\n")
	for lineno, eachLine := range lines {
		eachLine = strings.TrimSpace(eachLine)
		parts := strings.Split(eachLine, " ")
		if len(parts) != 2 {
			log.Println(fmt.Sprintf("incorrect number of parts on line %d: %#v", lineno, eachLine))
			continue
		}

		if result := rps.Play(parts[0], parts[1]); result > 0 {
			totalScore += result
		} else {
			panic(fmt.Sprintf("invalid score (%d) on line %d: %#v", result, lineno, eachLine))
		}
	}
	log.Printf("total score: %d", totalScore)
}

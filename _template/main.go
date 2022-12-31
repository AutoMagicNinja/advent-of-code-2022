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
	for lineNumber, eachLine := range lines {
	}
}

package main

import (
	"strings"

	"github.com/automagicninja/advent-of-code-2022/logger"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/afero"
	flag "github.com/spf13/pflag"
)

var (
	log           = logger.ResetLevel(zapcore.DebugLevel)
	inputFilename string
)

func init() {
	flag.StringVar(&inputFilename, "input", "", "Input filename")
}

func firstStar(lines []string) {
	var answer int
	for lineNumber, eachLine := range lines {
	}
	log.Infof("first star: %d", answer)
}

// func secondStar(lines []string) {
// 	var answer int
// 	for lineNumber, eachLine := range lines {
// 	}
// 	log.Infof("second star: %d", answer)
// }

func main() {
	var (
		err     error
		inputFs afero.Fs = afero.NewReadOnlyFs(afero.NewOsFs())
		rawData []byte
		lines   []string
	)

	flag.Parse()
	log.Debugf("input filename: %#v\n", inputFilename)

	rawData, err = afero.ReadFile(inputFs, inputFilename)
	if err != nil {
		panic(err)
	}

	lines = strings.Split(strings.TrimSpace(string(rawData)), "\n")
	firstStar(lines)
	// secondStar(lines)
}

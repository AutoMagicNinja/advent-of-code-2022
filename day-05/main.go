package main

import (
	"fmt"
	"sort"
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
	// log = logger.ResetLevel(zapcore.InfoLevel)
	var answer []rune
	drawing, instructions := parseInput(lines)
	// log.Debugf("drawing: %#v", drawing)
	fmt.Printf("drawing: %#v\n", drawing)
	stacks := CreateStacks(drawing)
	for k, v := range stacks {
		log.Debugf("stack id: %d, top: %c, size: %d", k, v.Peek(), v.Size())
	}
	Execute(&stacks, instructions)
	ids := make([]int, 0)
	for k, v := range stacks {
		log.Debugf("stack id: %d, top: %c, size: %d", k, v.Peek(), v.Size())
		ids = append(ids, k)
	}
	sort.Ints(ids)
	for _, id := range ids {
		s := stacks[id]
		answer = append(answer, s.Peek())
	}
	log.Infof("first star: %s", string(answer))
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

	lines = strings.Split(strings.TrimRight(string(rawData), "\r\n"), "\n")
	firstStar(lines)
	// secondStar(lines)
}

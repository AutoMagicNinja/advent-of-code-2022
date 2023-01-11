package main

import (
	"strings"
)

// var log = logger.ResetLevel(zapcore.DebugLevel)

func reverseSlice(s []string) []string {
	var (
		inputLen int = len(s)
		mid      int = inputLen / 2
	)
	for i := 0; i < mid; i++ {
		var j int = inputLen - i - 1

		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Parses the input to determine the drawing.
// first return value is the drawing
// second return value is the instructions
// Note: this assumes the input file is complete, and that a blank line is present to denote the separation between the drawing and the instructions
// Note: the drawing is reversed
func parseInput(lines []string) ([]string, []string) {
	var (
		drawing      []string
		instructions []string
	)
	for i, v := range lines {
		log.Debugf("line %d: %#v", i, v)
		// first blank line is the end of the drawing
		if len(strings.TrimSpace(v)) <= 0 {
			drawing = lines[:i]
			instructions = lines[i+1:]
			log.Debugf("drawing ends on line %d", i+1)
			break
		}
	}
	return reverseSlice(drawing), instructions
}

package main

import (
	"regexp"
	"strconv"
	"unicode"

	"github.com/automagicninja/advent-of-code-2022/day-05/stack"
)

// var log = logger.ResetLevel(zapcore.DebugLevel)

var (
	RE_STACK_IDENTIFIERS = regexp.MustCompile(`(?:\s*?(\d+)\s*?)`)
)

// drawing: line at idx 0 MUST be the one with the indices
func CreateStacks(drawing []string) map[int]*stack.Stack {
	stackIdsLine := []byte(drawing[0])
	stackIdIndicies := make(map[int]int, 0)
	stackIdMap := make(map[int]*stack.Stack, 0)

	allStackIdIndexLocations := RE_STACK_IDENTIFIERS.FindAllSubmatchIndex([]byte(stackIdsLine), -1)

	log.Debugf("all stack indices: %#v", allStackIdIndexLocations)
	for _, loc := range allStackIdIndexLocations {
		// these are all groups of four integers, with the following definitions:
		//  - group start
		//  - group end
		//  - match start
		//  - match end
		matchStart := loc[2]
		matchEnd := loc[3]
		idStr := string(stackIdsLine[matchStart:matchEnd])
		if len(idStr) > 1 {
			log.Warnf("id with length longer than 1: %#v", idStr)
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Warnf("could not parse %s as int")
		}
		log.Debugf("stack id %s starts at %d and ends at %d and was converted to %d", idStr, matchStart, matchEnd, id)

		// save the index of the stack id
		// all items on the stack in the drawing are at the same index as the stack ID
		// TODO: what if the stack ids are multiple wide?
		stackIdIndicies[id] = matchStart
		// make a new stack with the ID
		stackIdMap[id] = stack.New()
	}

	// at this point, we have all the stack IDs and their indices
	stackDrawing := drawing[1:]
	for _, eachLine := range stackDrawing {
		for sId, s := range stackIdMap {
			indexOfPossibleItem := stackIdIndicies[sId]
			if indexOfPossibleItem >= len(eachLine) {
				continue
			} else if item := rune(eachLine[indexOfPossibleItem]); unicode.IsLetter(item) {
				s.Push(item)
			}
		}
	}
	return stackIdMap
}

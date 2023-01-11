package main

import (
	"regexp"
	"strconv"

	"github.com/automagicninja/advent-of-code-2022/day-05/stack"
)

var (
	RE_INSTRUCTION = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
)

func mustConvertToInteger(s string) int {
	if value, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return value
	}
}

func execute(stacks *map[int]*stack.Stack, instruction string) {
	var (
		count, src, dst int
	)
	matches := RE_INSTRUCTION.FindAllStringSubmatch(instruction, -1)
	if len(matches) > 1 {
		log.Warnf("skipped instruction: `%s`; got more than one match for", instruction)
		return
	}
	instructionData := matches[0]
	if len(instructionData) != 4 {
		log.Warnf("skipped instruction: `%s`; got more 4 parts: %#v", instruction, instructionData)
		return
	}
	count = mustConvertToInteger(instructionData[1])
	src = mustConvertToInteger(instructionData[2])
	dst = mustConvertToInteger(instructionData[3])

	log.Infof("executing: %s", instruction)
	for i := 0; i < count; i++ {
		srcStack := (*stacks)[src]
		dstStack := (*stacks)[dst]
		log.Debugf("pre >> src: %d; size: %d; top: %c; contents: %#v", src, srcStack.Size(), srcStack.Peek(), srcStack.Contents())
		log.Debugf("pre >> dst: %d; size: %d; top: %c; contents: %#v", dst, dstStack.Size(), dstStack.Peek(), dstStack.Contents())
		dstStack.Push(srcStack.Pop())
		log.Debugf("post >> src: %d; size: %d; top: %c; contents: %#v", src, srcStack.Size(), srcStack.Peek(), srcStack.Contents())
		log.Debugf("post >> dst: %d; size: %d; top: %c; contents: %#v", dst, dstStack.Size(), dstStack.Peek(), dstStack.Contents())
	}
}

func Execute(stacks *map[int]*stack.Stack, instructions []string) {
	for _, v := range instructions {
		execute(stacks, v)
	}
}

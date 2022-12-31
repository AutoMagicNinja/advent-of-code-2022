package main

import (
	"strings"

	"github.com/automagicninja/advent-of-code-2022/day-03/set"
)

const SEARCH_ERROR = rune(0)

func findMisplaced(rucksack string) rune {
	var (
		items = []rune(strings.TrimSpace(rucksack))
		count = len(items)
	)

	if count%2 != 0 {
		return SEARCH_ERROR
	}

	firstCompartmentUnique := set.FromSlice(items[:count/2])
	secondCompartmentUnique := set.FromSlice(items[count/2:])

	diff := set.Union(firstCompartmentUnique, secondCompartmentUnique).ToSlice()
	if len(diff) != 1 {
		return SEARCH_ERROR
	}

	return diff[0]
}

func findBadge(elf1, elf2, elf3 string) rune {
	var (
		badge []rune
		union *set.Set
	)
	union = set.Union(
		set.FromSlice([]rune(elf1)),
		set.FromSlice([]rune(elf2)),
	)
	badge = set.Union(
		union,
		set.FromSlice([]rune(elf3)),
	).ToSlice()

	return badge[0]
}

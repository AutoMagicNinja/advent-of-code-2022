package main

const (
	LOWERCASE_LOWER_BOUND = rune('a')
	LOWERCASE_UPPER_BOUND = rune('z')
	UPPERCASE_LOWER_BOUND = rune('A')
	UPPERCASE_UPPER_BOUND = rune('Z')
)

func getValueOfItem(item rune) int {
	if (item >= LOWERCASE_LOWER_BOUND) && (item <= LOWERCASE_UPPER_BOUND) {
		return int(item-LOWERCASE_LOWER_BOUND) + 1
	} else if (item >= UPPERCASE_LOWER_BOUND) && (item <= UPPERCASE_UPPER_BOUND) {
		return int(item-UPPERCASE_LOWER_BOUND) + 27
	}
	return -1
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getValueOfItem(T *testing.T) {
	assert.Equal(T, 1, getValueOfItem(rune('a')))
	assert.Equal(T, 13, getValueOfItem(rune('m')))
	assert.Equal(T, 26, getValueOfItem(rune('z')))
	assert.Equal(T, 27, getValueOfItem(rune('A')))
	assert.Equal(T, 39, getValueOfItem(rune('M')))
	assert.Equal(T, 52, getValueOfItem(rune('Z')))
	assert.Equal(T, -1, getValueOfItem(rune(0x1795)))
}

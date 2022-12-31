package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMisplacedOK(T *testing.T) {
	assert.Equal(T, 'a', findMisplaced("abac"))
	assert.Equal(T, 'a', findMisplaced("abbbccdeefghaijklmnopqrs"))
}

func Test_findMisplacedNotOK(T *testing.T) {
	assert.Equal(T, SEARCH_ERROR, findMisplaced("abcd"))
	assert.Equal(T, SEARCH_ERROR, findMisplaced("aba"))
}

func Test_findMisplaceExapmle(T *testing.T) {
	assert.Equal(T, 'p', findMisplaced("vJrwpWtwJgWrhcsFMMfFFhFp"))
	assert.Equal(T, 'L', findMisplaced("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
}

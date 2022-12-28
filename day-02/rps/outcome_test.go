package rps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_determineYourMove(T *testing.T) {
	assert.Equal(T, Paper, determineYourMove(Rock, Win))
	assert.Equal(T, Scissors, determineYourMove(Paper, Win))
	assert.Equal(T, Rock, determineYourMove(Scissors, Win))

	assert.Equal(T, Rock, determineYourMove(Rock, Draw))
	assert.Equal(T, Paper, determineYourMove(Paper, Draw))
	assert.Equal(T, Scissors, determineYourMove(Scissors, Draw))

	assert.Equal(T, Scissors, determineYourMove(Rock, Loss))
	assert.Equal(T, Rock, determineYourMove(Paper, Loss))
	assert.Equal(T, Paper, determineYourMove(Scissors, Loss))
}

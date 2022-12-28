package rps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GameLogic(T *testing.T) {
	assert.Equal(T, Win, int(Result(Scissors, Rock)))
	assert.Equal(T, Win, int(Result(Rock, Paper)))
	assert.Equal(T, Win, int(Result(Paper, Scissors)))

	assert.Equal(T, Draw, int(Result(Rock, Rock)))
	assert.Equal(T, Draw, int(Result(Paper, Paper)))
	assert.Equal(T, Draw, int(Result(Scissors, Scissors)))

	assert.Equal(T, Loss, int(Result(Rock, Scissors)))
	assert.Equal(T, Loss, int(Result(Paper, Rock)))
	assert.Equal(T, Loss, int(Result(Scissors, Paper)))
}

func Test_ScoreLogic(T *testing.T) {
	assert.Equal(T, int(Win+Rock), Play("C", "X"))
	assert.Equal(T, int(Win+Paper), Play("A", "Y"))
	assert.Equal(T, int(Win+Scissors), Play("B", "Z"))

	assert.Equal(T, int(Draw+Rock), Play("A", "X"))
	assert.Equal(T, int(Draw+Paper), Play("B", "Y"))
	assert.Equal(T, int(Draw+Scissors), Play("C", "Z"))

	assert.Equal(T, int(Loss+Scissors), Play("A", "Z"))
	assert.Equal(T, int(Loss+Rock), Play("B", "X"))
	assert.Equal(T, int(Loss+Paper), Play("C", "Y"))
}

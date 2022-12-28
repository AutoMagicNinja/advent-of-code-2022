package rps

type Move int

const (
	Rock     Move = iota + 1 // Rock
	Paper                    // Paper
	Scissors                 // Scissors
)

type Outcome int

const (
	Loss Outcome = iota // Loss
	Draw Outcome = 3    // Draw
	Win  Outcome = 6    // Win
)

//go:generate stringer -type Move
//go:generate stringer -type Outcome

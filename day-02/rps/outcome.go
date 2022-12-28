package rps

import "log"

var OUTCOME_MAP map[string]Outcome = map[string]Outcome{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

func determineYourMove(opponents Move, outcome Outcome) Move {
	switch outcome {
	case Win:
		if opponents == Scissors {
			return Rock
		} else {
			return opponents + 1
		}
	case Draw:
		return opponents
	case Loss:
		if opponents == Rock {
			return Scissors
		} else {
			return opponents - 1
		}
	}
	return -1
}

func RequireOutcome(Opponent, DesiredOutcome string) int {
	var (
		outcome   Outcome
		opponents Move
		yours     Move
		isValid   bool
	)

	if outcome, isValid = OUTCOME_MAP[DesiredOutcome]; !isValid {
		return -1
	} else if opponents, isValid = OPPONENT_MAP[Opponent]; !isValid {
		return -1
	}

	yours = determineYourMove(opponents, outcome)

	log.Printf(">> DEBUG << Desired: %s, Opponent: %s, You: %s",
		outcome.String(), opponents.String(), yours.String())

	// poor-man's version of assert
	if outcome != Result(opponents, yours) {
		return -1
	}

	return Score(opponents, yours)
}

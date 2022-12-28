package rps

var OPPONENT_MAP map[string]Move = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var YOUR_MAP map[string]Move = map[string]Move{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// Result is from the perspective of the entity that is "You"
func Result(Opponent, You Move) Outcome {
	if You > Opponent {
		if (You == Scissors) && (Opponent == Rock) {
			return Loss
		}
		return Win
	} else if You == Opponent {
		return Draw
	} else {
		if (You == Rock) && (Opponent == Scissors) {
			return Win
		}
		return Loss
	}
}

func Score(Opponent, You Move) int {
	return int(Result(Opponent, You)) + int(You)
}

func Play(Opponent, You string) int {
	var (
		opponents Move
		yours     Move
		isValid   bool
	)

	if opponents, isValid = OPPONENT_MAP[Opponent]; !isValid {
		return -1
	}

	if yours, isValid = YOUR_MAP[You]; !isValid {
		return -1
	}

	return Score(opponents, yours)
}

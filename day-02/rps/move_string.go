// Code generated by "stringer -type Move"; DO NOT EDIT.

package rps

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Rock-1]
	_ = x[Paper-2]
	_ = x[Scissors-3]
}

const _Move_name = "RockPaperScissors"

var _Move_index = [...]uint8{0, 4, 9, 17}

func (i Move) String() string {
	i -= 1
	if i < 0 || i >= Move(len(_Move_index)-1) {
		return "Move(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Move_name[_Move_index[i]:_Move_index[i+1]]
}

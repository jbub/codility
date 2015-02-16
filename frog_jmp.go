package codility

// FrogJmp counts minimal number of jumps from position X to Y.
func FrogJmp(X int, Y int, D int) int {
	distance := Y - X
	jumps := distance / D
	remains := distance % D

	if remains > 0 {
		return jumps + 1
	}

	return jumps
}

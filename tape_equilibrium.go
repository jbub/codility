package codility

// TapeEquilibrium minimizes the value |(A[0] + ... + A[P-1]) - (A[P] + ... + A[N-1])|.
func TapeEquilibrium(A []int) int {
	N := len(A)
	s := sum(A)
	i := 0
	left := 0
	right := s
	min := 0
	set := false

	for i < N-1 {
		left += A[i]
		right -= A[i]
		diff := abs(left - right)

		if !set {
			min = diff
			set = true
		} else {
			if diff < min {
				min = diff
			}
		}

		i += 1
	}

	return min
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

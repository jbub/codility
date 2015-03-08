package permcheck

// Solution checks whether array A is a permutation.
func Solution(A []int) int {
	N := len(A)
	mapping := make(map[int]struct{})

	for _, v := range A {
		if v > N {
			return 0
		}
		mapping[v] = struct{}{}
	}

	if len(mapping) != N {
		return 0
	}

	return 1
}

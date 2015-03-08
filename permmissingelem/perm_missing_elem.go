package permmissingelem

// Solution finds the missing element in a given permutation.
func Solution(A []int) int {
	N := len(A)

	actual := sum(A)
	expected := sum(gen(N + 1))

	return expected - actual
}

func sum(A []int) int {
	s := 0
	for _, n := range A {
		s += n
	}
	return s
}

func gen(n int) []int {
	var arr []int
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	return arr
}

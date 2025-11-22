package kata

func CountBy(x, n int) []int {
	result := []int{}

	for i := x; i <= x*n; i += x {
		result = append(result, i)
	}

	return result
}

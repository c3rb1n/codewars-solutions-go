package kata

func monkeyCount(n int) []int {
	var result []int

	for i := 1; i <= n; i++ {
		result = append(result, i)
	}

	return result
}

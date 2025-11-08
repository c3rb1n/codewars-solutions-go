package kata

func CountSheeps(numbers []bool) int {
	result := 0

	for _, value := range numbers {
		if value {
			result++
		}
	}

	return result
}

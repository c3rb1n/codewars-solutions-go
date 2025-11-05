package kata

func Maps(x []int) []int {
	var result []int

	for _, value := range x {
		result = append(result, value*2)
	}

	return result
}

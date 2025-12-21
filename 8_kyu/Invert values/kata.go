package kata

func Invert(arr []int) []int {
	result := make([]int, 0, len(arr))

	for _, v := range arr {
		result = append(result, -v)
	}

	return result
}

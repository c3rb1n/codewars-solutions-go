package kata

func Grow(arr []int) int {
	res := 1

	for _, value := range arr {
		res *= value
	}

	return res
}

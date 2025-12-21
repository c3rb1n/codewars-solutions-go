package kata

func QuarterOf(month int) int {
	var result int

	switch {
	case month <= 3:
		result = 1
	case month <= 6:
		result = 2
	case month <= 9:
		result = 3
	default:
		result = 4
	}

	return result
}

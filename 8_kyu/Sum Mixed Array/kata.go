package kata

import "strconv"

func SumMix(arr []any) int {
	result := 0

	for _, v := range arr {
		switch v := v.(type) {
		case int:
			result += v
		case string:
			vInt, _ := strconv.Atoi(v)
			result += vInt
		}
	}

	return result
}

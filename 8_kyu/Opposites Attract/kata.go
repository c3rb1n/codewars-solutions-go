package kata

func LoveFunc(flower1, flower2 int) bool {
	case1 := flower1%2 == 0 && flower2%2 != 0
	case2 := flower2%2 == 0 && flower1%2 != 0

	return case1 || case2
}

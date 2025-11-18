package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	upperStr := strings.ToUpper(str)
	reversedStr := ""

	for i := len(upperStr) - 1; i >= 0; i-- {
		reversedStr += string(upperStr[i])
	}

	return reversedStr == upperStr
}

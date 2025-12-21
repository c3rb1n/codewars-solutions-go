package kata

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	strElements := strings.Split(str, " ")
	strLen := len(strElements)
	result := ""

	for i := strLen - 1; i >= 0; i-- {
		result += fmt.Sprint(strElements[i], " ")
	}

	return strings.TrimSpace(result)
}

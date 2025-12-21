package kata

func RemoveChar(word string) string {
	wordLength := len(word)

	if wordLength == 2 {
		return ""
	}

	result := ""

	for i, v := range word {
		if i > 0 && i < wordLength-1 {
			result += string(v)
		}
	}

	return result
}

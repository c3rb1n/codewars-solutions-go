package kata

func BoolToWord(word bool) string {
	var result string

	if word {
		result = "Yes"
	} else {
		result = "No"
	}

	return result
}

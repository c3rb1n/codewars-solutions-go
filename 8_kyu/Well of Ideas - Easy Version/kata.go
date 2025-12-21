package kata

func Well(x []string) string {
	result := 0

	for _, value := range x {
		if value == "good" {
			result++
		}
	}

	if result > 2 {
		return "I smell a series!"
	} else if result > 0 {
		return "Publish!"
	}

	return "Fail!"
}

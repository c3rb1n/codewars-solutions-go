package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	positivesCounter := 0
	sumOfNegatives := 0

	if len(numbers) != 0 && numbers != nil {
		for _, value := range numbers {
			if value > 0 {
				positivesCounter++
			} else if value < 0 {
				sumOfNegatives += value
			}
		}

		res = append(res, positivesCounter)
		res = append(res, sumOfNegatives)
	}

	return res
}

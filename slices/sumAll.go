package slices

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumber := len(numbersToSum)
	sums := make([]int, lengthOfNumber)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

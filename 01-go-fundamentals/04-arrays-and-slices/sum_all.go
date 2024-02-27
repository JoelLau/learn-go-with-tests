package arraysandslices

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// using `make` seems to be almost twice as fast!
func SumAll2(numbersToSum ...[]int) []int {
	// initialise slice with specific length
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

package arraysandslices

// a "tail" is the sum of all items in the slice, excluding
// the first item.
func SumAllTails(numbersToSum ...[]int) (sum []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) <= 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers[1:]))
		}
	}
	return
}

func SumAll(numbers ...[]int) (sums []int) {
	for i := 0; i < len(numbers); i++ {
		sums = append(sums, Sum(numbers[i]))
	}
	return
}

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

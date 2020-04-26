package array

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTrails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sum := 0
		if len(numbers) > 1 {
			sum = Sum(numbers[1:])
		}
		sums = append(sums, sum)
	}
	return
}

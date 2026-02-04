package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// // (0, 1)
	// // ([]int{1, 2}, []int{0, 9})
	// // i : 0, 1 | numbers := [1, 2], [0, 9]
	// for i, numbers := range numbersToSum {
	// 	fmt.Println(i)
	// 	sums[i] = Sum(numbers)
	// }

	return sums

}

package calc

func IsBiggerThan(a int, b int) bool {
	return a > b
}

func Average(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum / len(numbers)
}

package calc

// Returns the result is bigger than second number
func IsBiggerThan(a int, b int) bool {
	return a > b
}

// Returns the average of given list numbers
func Average(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum / len(numbers)
}

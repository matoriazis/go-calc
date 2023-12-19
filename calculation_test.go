package calc

import "testing"

func TestAverage(t *testing.T) {
	result := Average([]int{5, 4, 6})
	if result != 5 {
		t.Errorf("Average(3, 4) = %d; want 5", result)
	}
}

func TestIsBiggerThan(t *testing.T) {
	result := IsBiggerThan(34, 23)
	if result != true {
		t.Errorf("IsBiggerThan(34, 23); want true")
	}
}

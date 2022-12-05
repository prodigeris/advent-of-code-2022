package day_4_camp_cleanup

import (
	"reflect"
	"testing"
)

func TestSumFullyContain(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	actual, _ := sum(input)
	expected := 2
	if actual != expected {
		t.Errorf("Expected Int(%d) is not same as"+
			" actual int (%d)", expected, actual)
	}
}

func TestSumOverlap(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	_, actual := sum(input)
	expected := 4
	if actual != expected {
		t.Errorf("Expected Int(%d) is not same as"+
			" actual int (%d)", expected, actual)
	}
}

func TestReturnStringRangeToSlice(t *testing.T) {
	input := map[string]map[int]int{
		"2-4": {2: 2, 3: 3, 4: 4},
		"2-3": {2: 2, 3: 3},
		"5-7": {5: 5, 6: 6, 7: 7},
		"2-8": {2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8},
		"6-6": {6: 6},
		"2-6": {2: 2, 3: 3, 4: 4, 5: 5, 6: 6},
	}

	for str, expected := range input {
		actual := stringRangesToSlice(str)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected []Int(%d) is not same as"+
				" actual []int (%d)", expected, actual)
		}
	}
}

package day_3_rucksack

import "testing"

func TestReturnGeeks(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	actual := calculateTotal(input)
	expected := 157
	if actual != expected {
		t.Errorf("Expected Int(%d) is not same as"+
			" actual int (%d)", expected, actual)
	}
}

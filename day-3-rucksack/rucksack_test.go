package day_3_rucksack

import "testing"

func TestReturnCorrectTotalPriority(t *testing.T) {
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

func TestReturnCorrectPriorityForLetter(t *testing.T) {
	input := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'A': 27,
		'Z': 52,
	}
	for letter, expected := range input {
		actual := getPriority(letter)
		if actual != expected {
			t.Errorf("Expected Int(%d) is not same as"+
				" actual int (%d) for letter (%s)", expected, actual, string(letter))
		}
	}
}

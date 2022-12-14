package day_5_supply_stack

import (
	"reflect"
	"testing"
)

func TestParsesInputCorrectly(t *testing.T) {
	input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	expectedStacks := map[int][]string{1: {"N", "Z"}, 2: {"D", "C", "M"}, 3: {"P"}}
	expectedActions := []action{{1, 2, 1}, {3, 1, 3}, {2, 2, 1}, {1, 1, 2}}

	stacks, actions := parse(input)

	if !reflect.DeepEqual(expectedStacks, stacks) {
		t.Errorf("Expected (%v) is not same as"+
			" actual (%v)", expectedStacks, stacks)
	}

	if !reflect.DeepEqual(expectedActions, actions) {
		t.Errorf("Expected []action(%v) is not same as"+
			" actual []action (%v)", expectedActions, actions)
	}

}

func TestParse5Stacks(t *testing.T) {
	input := []string{
		"    [D]            ",
		"[N] [C]            ",
		"[Z] [M] [P] [A] [X]",
		" 1   2   3   4   5 "}

	actual := parseStacks(input)
	expected := map[int][]string{1: {"N", "Z"}, 2: {"D", "C", "M"}, 3: {"P"}, 4: {"A"}, 5: {"X"}}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected (%v) is not same as"+
			" actual (%v)", expected, actual)
	}
}

func TestCrateMover9000(t *testing.T) {
	input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	actual := calc(input, createMover9000)
	expected := "CMZ"
	if actual != expected {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expected, actual)
	}
}

func TestCrateMover9001(t *testing.T) {
	input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	actual := calc(input, createMover9001)
	expected := "MCD"
	if actual != expected {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expected, actual)
	}
}

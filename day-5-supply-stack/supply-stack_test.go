package day_5_supply_stack

import (
	"reflect"
	"testing"
)

func TestParsesInputCorrectly(t *testing.T) {
	input := []string{"[D]",
		"[N] [C]",
		"[Z] [M] [P]",
		"1   2   3",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	expectedStacks := [][]rune{{'N', 'Z'}, {'D', 'C', 'M'}, {'P'}}
	expectedActions := []action{{1, 2, 1}, {3, 1, 3}, {2, 2, 1}, {1, 1, 2}}

	stacks, actions := parse(input)

	if !reflect.DeepEqual(expectedStacks, stacks) {
		t.Errorf("Expected [][]rune(%v) is not same as"+
			" actual [][]rune (%v)", expectedStacks, stacks)
	}

	if !reflect.DeepEqual(expectedActions, actions) {
		t.Errorf("Expected []action(%v) is not same as"+
			" actual []action (%v)", expectedActions, actions)
	}

}

//func TestSumFullyContain(t *testing.T) {
//	input := []string{
//		"[D]",
//		"[N] [C]",
//		"[Z] [M] [P]",
//	}
//
//	input2 := []string{
//		"move 1 from 2 to 1",
//		"move 3 from 1 to 3",
//		"move 2 from 2 to 1",
//		"move 1 from 1 to 2",
//	}
//
//	actual := calc(input, input2)
//	expected := "CMZ"
//	if actual != expected {
//		t.Errorf("Expected String(%s) is not same as"+
//			" actual string (%s)", expected, actual)
//	}
//}

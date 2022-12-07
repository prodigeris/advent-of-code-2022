package day_5_supply_stack

import "testing"

func TestSumFullyContain(t *testing.T) {
	input := []string{
		"[D]",
		"[N] [C]",
		"[Z] [M] [P]",
	}

	input2 := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	actual := calc(input, input2)
	expected := "CMZ"
	if actual != expected {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expected, actual)
	}
}

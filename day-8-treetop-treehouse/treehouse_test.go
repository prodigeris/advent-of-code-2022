package day_8_treetop_treehouse

import (
	"reflect"
	"testing"
)

func TestCountVisible(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	actual := countVisible(input)
	if actual != 21 {
		t.Errorf("Expected int(%d) is not same as"+
			" actual int(%d)", 21, actual)
	}
}

func TestGetTop(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expected := []int{0, 5, 5}

	actual := getTop(input, 3, 1)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected int(%d) is not same as"+
			" actual int(%d)", expected, actual)
	}
}

func TestGetLeft(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expected := []int{3, 3}

	actual := getLeft(input, 3, 2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected int(%d) is not same as"+
			" actual int(%d)", expected, actual)
	}
}

func TestGetRight(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expected := []int{4, 9}

	actual := getRight(input, 3, 2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected int(%d) is not same as"+
			" actual int(%d)", expected, actual)
	}
}

func TestGetBottom(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	expected := []int{5, 3, 5, 3}

	actual := getBottom(input, 0, 2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected int(%d) is not same as"+
			" actual int(%d)", expected, actual)
	}
}

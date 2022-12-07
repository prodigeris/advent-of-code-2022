package day_7_no_space_left

import "testing"

func TestCrateMover9001(t *testing.T) {
	input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	actual := interpret(input)
	if actual != 95437 {
		t.Errorf("Expected int(%d) is not same as"+
			" actual int(%d)", 95437, actual)
	}
}

package day_8_treetop_treehouse

import "strconv"

func getTop(lines []string, r int, c int) []int {
	t := make([]int, 0)
	for i := 0; i < r; i++ {
		t = append(t, runeToInt(lines[i][c]))
	}

	return t
}

func getLeft(lines []string, r int, c int) []int {
	t := make([]int, 0)
	for i := 0; i < c; i++ {
		t = append(t, runeToInt(lines[r][i]))
	}

	return t
}

func getRight(lines []string, r int, c int) []int {
	t := make([]int, 0)
	for i := c + 1; i < len(lines[r]); i++ {
		t = append(t, runeToInt(lines[r][i]))
	}

	return t
}

func getBottom(lines []string, r int, c int) []int {
	t := make([]int, 0)
	for i := r + 1; i < len(lines); i++ {
		t = append(t, runeToInt(lines[i][c]))
	}

	return t
}

func runeToInt(r uint8) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		panic(err)
	}
	return i
}

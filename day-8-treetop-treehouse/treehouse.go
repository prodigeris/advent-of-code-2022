package day_8_treetop_treehouse

import (
	"advent-of-code-2022/common"
)

func Run() int {
	lines := common.ReadToLines("day-8-treetop-treehouse/input.txt")
	return countVisible(lines)
}

func countVisible(lines []string) int {
	t := 0
	for i, line := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}
		for j, l := range line {
			if j == 0 || j == len(line)-1 {
				continue
			}
			if isVisibleInlines(l, lines, i, j) {
				t++
			}
		}
	}

	return t + len(lines)*2 + len(lines[0])*2 - 4
}

func isVisible(c int, l []int, t []int, b []int, r []int) bool {
	return isVisibleInGroup(c, l) || isVisibleInGroup(c, t) || isVisibleInGroup(c, b) || isVisibleInGroup(c, r)
}

func isVisibleInGroup(c int, g []int) bool {
	for _, v := range g {
		if c <= v {
			return false
		}
	}
	return true
}

func isVisibleInlines(c int32, l []string, i int, j int) bool {
	return isVisible(runeToInt(uint8(c)), getLeft(l, i, j), getTop(l, i, j), getBottom(l, i, j), getRight(l, i, j))
}

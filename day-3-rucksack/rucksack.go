package day_3_rucksack

import (
	"advent-of-code-2022/common"
	"fmt"
)

func Run() {
	lines := common.ReadToLines("day-3-rucksack/input.txt")
	calculateTotal(lines)
}

func calculateTotal(lines []string) int {
	t := 0
	for _, line := range lines {
		fmt.Println(line)

		break
	}
	return t
}

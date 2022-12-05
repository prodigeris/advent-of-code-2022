package day2RockPaperScissors

import (
	"advent-of-code-2022/common"
	"strings"
)

type getScore func(s1 string, s2 string) int

func CalculateScore(s getScore) int {
	t := 0
	lines := common.ReadToLines("day-2-rock-paper-scissors/input.txt")
	for _, v := range lines {
		cols := strings.Split(v, " ")
		t += s(cols[0], cols[1])
	}
	return t
}

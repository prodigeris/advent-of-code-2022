package day2RockPaperScissors

import (
	"advent-of-code-2022/common"
	"fmt"
	"strings"
)

func Run() {
	t := 0
	lines := common.ReadToLines("day-2-rock-paper-scissors/input.txt")
	for _, v := range lines {
		actions := strings.Split(v, " ")
		t += calcScore(actions[0], actions[1])
	}
	fmt.Println(t)
}

// A, X - rock - 1
// B, Y - paper - 2
// C, Z - scissors - 3
func calcScore(a1 string, a2 string) int {
	results := map[string]int{
		"AY": 8,
		"AX": 4,
		"AZ": 3,
		"BY": 5,
		"BX": 1,
		"BZ": 9,
		"CY": 2,
		"CX": 7,
		"CZ": 6,
	}

	s, ok := results[a1+a2]
	if !ok {
		panic("Action is not allowed")
	}

	return s
}

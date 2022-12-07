package day_5_supply_stack

import (
	"advent-of-code-2022/common"
)

type action struct {
	amount int
	from   int
	to     int
}

type crateMover func(stacks *map[int][]string, action action)

func Run() string {
	lines := common.ReadToLines("day-5-supply-stack/input.txt")
	return calc(lines, createMover9001)
}

func calc(i []string, mover crateMover) string {
	stacks, actions := parse(i)

	for _, a := range actions {
		mover(&stacks, a)
	}

	s := ""
	for i := 1; i <= len(stacks); i++ {
		s += stacks[i][0]
	}

	return s
}

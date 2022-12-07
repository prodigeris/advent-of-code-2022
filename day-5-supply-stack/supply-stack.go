package day_5_supply_stack

import (
	"advent-of-code-2022/common"
)

type action struct {
	amount int
	from   int
	to     int
}

func Run() string {
	lines := common.ReadToLines("day-5-supply-stack/input.txt")
	return calc(lines)
}

func calc(i []string) string {
	stacks, actions := parse(i)

	for _, a := range actions {
		for i := 0; i < a.amount; i++ {
			stacks[a.to] = append(
				[]string{stacks[a.from][0]},
				stacks[a.to]...)
			stacks[a.from] = stacks[a.from][1:]
		}
	}

	s := ""
	for i := 1; i <= len(stacks); i++ {
		s += stacks[i][0]
	}

	return s
}

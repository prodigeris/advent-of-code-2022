package day_5_supply_stack

import "fmt"

type action struct {
	amount int
	from   int
	to     int
}

func Run() {
}

func calc(i []string) string {
	stacks, actions := parse(i)

	for _, a := range actions {
		for i := 0; i < a.amount; i++ {
			stacks[a.to] = append([]string{stacks[a.from][0]}, stacks[a.to]...)
			stacks[a.from] = stacks[a.from][1:]
		}
	}

	fmt.Println(stacks)

	s := ""
	for i := 1; i <= len(stacks); i++ {
		s += stacks[i][0]
	}

	return s
}

package day_5_supply_stack

func createMover9001(stacks *map[int][]string, a action) {
	l := make([]string, a.amount)
	copy(l, (*stacks)[a.from][0:a.amount])

	(*stacks)[a.to] = append(l, (*stacks)[a.to]...)
	(*stacks)[a.from] = (*stacks)[a.from][a.amount:]
}

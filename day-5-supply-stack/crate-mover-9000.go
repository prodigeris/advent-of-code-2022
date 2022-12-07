package day_5_supply_stack

func createMover9000(stacks *map[int][]string, a action) {
	for i := 0; i < a.amount; i++ {
		(*stacks)[a.to] = append(
			[]string{(*stacks)[a.from][0]},
			(*stacks)[a.to]...)
		(*stacks)[a.from] = (*stacks)[a.from][1:]
	}
}

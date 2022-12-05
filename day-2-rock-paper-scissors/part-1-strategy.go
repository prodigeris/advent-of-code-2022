package day2RockPaperScissors

func Part1Strategy(a1 string, a2 string) int {
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

package day2RockPaperScissors

func Part2Strategy(h string, r string) int {
	results := map[string]int{
		"AY": 4,
		"AX": 3,
		"AZ": 8,
		"BY": 5,
		"BX": 1,
		"BZ": 9,
		"CY": 6,
		"CX": 2,
		"CZ": 7,
	}

	s, ok := results[h+r]
	if !ok {
		panic("Action is not allowed")
	}

	return s
}

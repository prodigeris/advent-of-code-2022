package day_6_tuning_trouble

import (
	"advent-of-code-2022/common"
)

func Run() (int, int) {
	lines := common.ReadToLines("day-6-tuning-trouble/input.txt")
	return detect(lines[0], 4), detect(lines[0], 14)
}

func detect(s string, count int) int {
	for i := count; i <= len(s); i++ {
		if checkIfAllLettersDifferent(s[i-count : i]) {
			return i
		}
	}

	return 0
}

func checkIfAllLettersDifferent(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

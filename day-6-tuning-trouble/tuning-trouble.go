package day_6_tuning_trouble

import "fmt"

func detect(s string) int {
	for i := 4; i <= len(s); i++ {
		if checkIfAllLettersDifferent(s[i-4 : i]) {
			fmt.Println(s, i, i+4, s[i-4:i])
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

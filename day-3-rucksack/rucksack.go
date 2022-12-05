package day_3_rucksack

import (
	"advent-of-code-2022/common"
	"strings"
)

func Run() (int, int) {
	lines := common.ReadToLines("day-3-rucksack/input.txt")
	return calcByCompartments(lines), calcByGroups(lines)
}

func calcByCompartments(lines []string) int {
	t := 0
	for _, line := range lines {
		h := len(line) / 2
		for _, v := range line[:h] {
			if strings.ContainsRune(line[h:], v) {
				t += getPriority(v)
				break
			}
		}
	}
	return t
}

func calcByGroups(lines []string) int {
	t := 0
	for i := 0; i < len(lines); i += 3 {
		for _, v := range lines[i] {
			if strings.ContainsRune(lines[i+1], v) && strings.ContainsRune(lines[i+2], v) {
				t += getPriority(v)
				break
			}
		}
	}
	return t
}

func getPriority(v rune) int {
	scores := map[rune]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13,
		'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
		'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39,
		'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52}
	t, ok := scores[v]
	if !ok {
		panic("Letter not found")
	}
	return t
}

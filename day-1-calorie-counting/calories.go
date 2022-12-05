package calories

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	lines := readInput()
	elves := make([]int, 1)
	c := 0
	for _, v := range lines {
		if v == "" {
			c++
			elves = append(elves, 0)
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		elves[c] += i
	}

	sort.Ints(elves)
	top3 := elves[len(elves)-3:]

	fmt.Println("Top 3 elves:", top3)
	fmt.Println("Top 3 total calories:", sum(top3))
}

func sum(r []int) int {
	t := 0
	for _, v := range r {
		t += v
	}
	return t
}

func readInput() []string {
	d, err := os.ReadFile("day-1-calorie-counting/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(d), "\n")
}

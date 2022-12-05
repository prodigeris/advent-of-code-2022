package calories

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calories uint

func Run() {
	lines := readInput()
	elves := make(map[int]Calories)
	c := 0
	max := Calories(0)
	for _, v := range lines {
		if v == "" {
			c++
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		elves[c] += Calories(i)

		if elves[c] > max {
			max = elves[c]
		}
	}

	fmt.Println(max)
}

func readInput() []string {
	d, err := os.ReadFile("day-1-calorie-counting/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(d), "\n")
}

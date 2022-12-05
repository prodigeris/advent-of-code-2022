package day_4_camp_cleanup

import (
	"advent-of-code-2022/common"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	lines := common.ReadToLines("day-4-camp-cleanup/input.txt")
	fmt.Println(sum(lines))
}

func sum(lines []string) int {
	return len(lines)
}

func stringRangesToSlice(r string) []int {
	fromTo := strings.Split(r, "-")
	from, err1 := strconv.Atoi(fromTo[0])
	to, err2 := strconv.Atoi(fromTo[1])
	if err1 != nil || err2 != nil {
		panic("Invalid from to range")
	}

	ranges := make([]int, 0)
	for i := from; i <= to; i++ {
		ranges = append(ranges, i)
	}
	return ranges
}

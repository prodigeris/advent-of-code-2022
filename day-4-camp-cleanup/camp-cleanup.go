package day_4_camp_cleanup

import (
	"advent-of-code-2022/common"
	"strconv"
	"strings"
)

func Run() int {
	lines := common.ReadToLines("day-4-camp-cleanup/input.txt")
	return sum(lines)
}

func sum(lines []string) int {
	t := 0
	for _, v := range lines {
		ranges := strings.Split(v, ",")
		if doRangesOverlap(stringRangesToSlice(ranges[0]), stringRangesToSlice(ranges[1])) {
			t++
		}
	}

	return t
}

func stringRangesToSlice(r string) map[int]int {
	fromTo := strings.Split(r, "-")
	from, err1 := strconv.Atoi(fromTo[0])
	to, err2 := strconv.Atoi(fromTo[1])
	if err1 != nil || err2 != nil {
		panic("Invalid from to range")
	}

	ranges := make(map[int]int)
	for i := from; i <= to; i++ {
		ranges[i] = i
	}
	return ranges
}

func doRangesOverlap(r map[int]int, r2 map[int]int) bool {
	if len(r) > len(r2) {
		return doRangesOverlap(r2, r)
	}
	for _, v := range r {
		_, found := r2[v]
		if !found {
			return false
		}
	}
	return true
}

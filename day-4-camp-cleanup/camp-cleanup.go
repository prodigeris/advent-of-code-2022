package day_4_camp_cleanup

import (
	"advent-of-code-2022/common"
	"strconv"
	"strings"
)

func Run() (int, int) {
	lines := common.ReadToLines("day-4-camp-cleanup/input.txt")
	return sum(lines)
}

func sum(lines []string) (int, int) {
	c, o := 0, 0
	for _, v := range lines {
		ranges := strings.Split(v, ",")
		if doRangesFullyContain(stringRangesToSlice(ranges[0]), stringRangesToSlice(ranges[1])) {
			c++
		}
		if doRangesOverlap(stringRangesToSlice(ranges[0]), stringRangesToSlice(ranges[1])) {
			o++
		}
	}

	return c, o
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

func doRangesFullyContain(r map[int]int, r2 map[int]int) bool {
	if len(r) > len(r2) {
		return doRangesFullyContain(r2, r)
	}
	for _, v := range r {
		_, found := r2[v]
		if !found {
			return false
		}
	}
	return true
}

func doRangesOverlap(r map[int]int, r2 map[int]int) bool {
	for _, v := range r {
		_, found := r2[v]
		if found {
			return true
		}
	}
	return false
}

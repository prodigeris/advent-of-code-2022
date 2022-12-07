package day_5_supply_stack

import (
	"fmt"
	"regexp"
	"strconv"
)

func parse(l []string) (map[int][]string, []action) {
	stacksInput := make([]string, 0)
	actionsInput := make([]string, 0)
	isActions := false
	for _, v := range l {
		if v == "" {
			isActions = true
			continue
		}
		if isActions {
			actionsInput = append(actionsInput, v)
		} else {
			stacksInput = append(stacksInput, v)
		}
	}

	return parseStacks(stacksInput), parseActions(actionsInput)
}

func parseStacks(i []string) map[int][]string {
	s := make(map[int][]string)
	for k, v := range i {
		if k+1 == len(i) {
			break
		}
		for i := 0; i < len(v); i += 4 {
			re := regexp.MustCompile(`(\[|\]|\s|\d)`)
			l := re.ReplaceAllString(v[i:i+3], "")
			if l == "" {
				continue
			}

			k := i/4 + 1
			if _, ok := s[k]; !ok {
				s[k] = []string{}
			}
			s[k] = append(s[k], l)
		}
	}
	return s
}

func parseActions(a []string) (r []action) {
	for _, v := range a {
		re := regexp.MustCompile(`move (.+) from (.+) to (.+)`)
		match := re.FindStringSubmatch(v)
		fmt.Println(v, match)
		amount, err1 := strconv.Atoi(match[1])
		from, err2 := strconv.Atoi(match[2])
		to, err3 := strconv.Atoi(match[3])

		if err1 != nil || err2 != nil || err3 != nil {
			panic("Cannot convert string to int")
		}

		r = append(r, action{amount, from, to})
	}
	return
}

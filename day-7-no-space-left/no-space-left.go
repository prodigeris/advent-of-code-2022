package day_7_no_space_left

import (
	"advent-of-code-2022/common"
	"regexp"
	"strconv"
	"strings"
)

func Run() uint64 {
	lines := common.ReadToLines("day-7-no-space-left/input.txt")
	return interpret(lines)
}

type directory struct {
	path   string
	size   uint64
	parent *directory
}

func interpret(c []string) uint64 {
	d := directory{"home", 0, nil}
	current := &d
	directories := []*directory{&d}
	for i := 0; i < len(c); i++ {
		s := c[i]
		if s[0] == '$' {
			if s != "$ ls" {
				re, err := regexp.Compile(`cd (.*)`)
				if err != nil {
					panic(err)
				}
				match := re.FindStringSubmatch(s)
				if match[1] == "/" {
					continue
				}

				if match[1] == ".." {
					current = current.parent
				} else {
					newDir := directory{current.path + "/" + match[1], 0, current}
					current = &newDir
					directories = append(directories, &newDir)
				}
			}
		} else {
			if s[0] == 'd' {
				continue
			}

			sizes := strings.Split(c[i], " ")
			size, err := strconv.Atoi(sizes[0])
			if err != nil {
				panic(err)
			}
			addSize(current, size)
		}
	}

	t := uint64(0)
	for _, dir := range directories {
		if dir.size <= 100000 {
			t += dir.size
		}
	}

	return t
}

func addSize(current *directory, size int) {
	c := current
	for {
		c.size += uint64(size)
		if c.parent == nil {
			break
		}
		c = c.parent
	}
}

package common

import (
	"os"
	"strings"
)

func ReadToLines(path string) []string {
	d, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(d), "\n")
}

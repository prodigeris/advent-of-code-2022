package day_5_supply_stack

import "fmt"

type action struct {
	amount int
	from   int
	to     int
}

func Run() {

}

func parse(l []string) ([][]rune, []action) {
	for _, v := range l {
		fmt.Println(v)
	}

	return [][]rune{}, []action{}
}

func calc(i []string, i2 []string) string {
	return ""
}

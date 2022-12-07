package day_6_tuning_trouble

import "testing"

func TestCrateMover9001(t *testing.T) {
	input := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for i, r := range input {
		actual := detect(i)
		if actual != r {
			t.Errorf("Expected int(%d) is not same as"+
				" actual int(%d) %s", r, actual, i)
		}
	}

}

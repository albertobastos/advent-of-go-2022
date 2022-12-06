package main

import "testing"

func _Test(input string, part1exp int, part2exp int, t *testing.T) {
	part1, part2 := run(input)
	if part1 != part1exp {
		t.Errorf("part1 is wrong, expected %d but got %d", part1exp, part1)
	}
	if part2 != part2exp {
		t.Errorf("part2 is wrong, expected %d but got %d", part2exp, part2)
	}
}

func TestDemo(t *testing.T) {
	_Test("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19, t)
	_Test("bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23, t)
	_Test("nppdvjthqldpwncqszvftbrmjlhg", 6, 23, t)
	_Test("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29, t)
	_Test("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26, t)
}

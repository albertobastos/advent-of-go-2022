package main

import "testing"

func _Test(file string, row int, distressmax int, part1exp int, part2exp int, t *testing.T) {
	part1, part2 := run(row, distressmax, file)
	if part1 != part1exp {
		t.Errorf("part1 is wrong, expected %d but got %d", part1exp, part1)
	}
	if part2 != part2exp {
		t.Errorf("part2 is wrong, expected %d but got %d", part2exp, part2)
	}
}

func TestDemo(t *testing.T) {
	_Test("demo.txt", 10, 20, 26, 56000011, t)
}

package main

import "testing"

func TestDemo(t *testing.T) {
	part1exp := 2
	part2exp := 4
	part1, part2 := run("demo.txt")
	if part1 != part1exp {
		t.Errorf("part1 is wrong, expected %d but got %d", part1exp, part1)
	}
	if part2 != part2exp {
		t.Errorf("part2 is wrong, expected %d but got %d", part2exp, part2)
	}
}

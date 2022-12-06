package main

import "testing"

func TestDemo(t *testing.T) {
	part1exp := "CMZ"
	part2exp := "MCD"
	part1, part2 := run("demo.txt")
	if part1 != part1exp {
		t.Errorf("part1 is wrong, expected %s but got %s", part1exp, part1)
	}
	if part2 != part2exp {
		t.Errorf("part2 is wrong, expected %s but got %s", part2exp, part2)
	}
}

package main

import "testing"

func _Test(file string, part1exp int, part2exp []string, t *testing.T) {
	part1, part2 := run(file)
	if part1 != part1exp {
		t.Errorf("part1 is wrong, expected %d but got %d", part1exp, part1)
	}
	if len(part2) != len(part2exp) {
		t.Errorf("part2 is wrong, expected length %d but got %d", len(part2exp), len(part2))
	}
	for i, row := range part2 {
		if row != part2exp[i] {
			t.Errorf("part2 row %d is wrong, expected %s but got %s", i, part2exp[i], row)
		}
	}
}

func TestDemo(t *testing.T) {
	display := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	_Test("demo.txt", 13140, display, t)
}

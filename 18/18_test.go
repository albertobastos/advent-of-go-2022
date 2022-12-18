package main

import "testing"

func _Test(file string, part1exp int, part2exp int, t *testing.T) {
	part1, part2 := run(file)
	if part1 != part1exp {
		t.Errorf("file %s part1 is wrong, expected %d but got %d", file, part1exp, part1)
	}
	if part2 != part2exp {
		t.Errorf("file %s part2 is wrong, expected %d but got %d", file, part2exp, part2)
	}
}

func TestDemo(t *testing.T) {
	_Test("demo.txt", 64, 58, t)
	/*_Test("demo2.txt", 10, 10, t)
	_Test("demo_3x3_full.txt", 54, 54, t)
	_Test("demo_3x3_1bubble.txt", 60, 54, t)*/
}

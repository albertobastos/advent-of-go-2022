package main

import "fmt"

type Priority uint8

func run(file string) (int, int) {
	return doPart1(file), doPart2(file)
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func priority(c rune) Priority {
	// unicodes: a..z = 97..122, A..Z = 65..90
	// priorities: a..z = 1..26, A..Z = 27..52
	if c >= 97 && c <= 122 {
		return Priority(c - 96)
	} else {
		return Priority(c - 38)
	}
}

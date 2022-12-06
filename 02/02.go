package main

import "fmt"

const (
	ROCK     int = 1
	PAPER        = 2
	SCISSORS     = 3
	INVALID      = -1
)

const (
	LOSE int = 0
	DRAW     = 3
	WIN      = 6
)

type Play struct {
	opponent int
	player   int
	result   int
}

func run(file string) (int, int) {
	return doPart1(file), doPart2(file)
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func readOpponent(c byte) int {
	if c == 'A' {
		return ROCK
	} else if c == 'B' {
		return PAPER
	} else if c == 'C' {
		return SCISSORS
	}
	return INVALID
}

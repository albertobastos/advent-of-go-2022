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

func main() {
	fmt.Println("Part 1 =", doPart1())
	fmt.Println("Part 2 =", doPart2())
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

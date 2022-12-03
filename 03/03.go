package main

import "fmt"

type Priority uint8

func main() {
	fmt.Println("Part1 =", doPart1())
	fmt.Println("Part2 =", doPart2())
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

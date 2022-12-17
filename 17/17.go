package main

import (
	"fmt"
	"io/ioutil"
)

const ROUNDS_PART1 = 2022
const ROUNDS_PART2 = 1000000000000

const WIDTH = 7
const LEFT = -1
const RIGHT = 1

type Jets []int
type Floor [WIDTH]int
type Point [2]int // x, y
type Rock []Point

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	jets := readFile(file)
	part1 := playRounds(jets, ROUNDS_PART1)
	fmt.Println("jets:", len(jets))
	// runs forever
	part2 := -1 // playRounds(jets, ROUNDS_PART2)
	return part1, part2
}

func playRounds(jets Jets, rounds int) int {
	floor := Floor{}
	stall := []Point{}
	ji := -1
	for ri := 0; ri < rounds; ri++ {
		r := newRock(ri, floor)
		//fmt.Println("round", i, "of", rounds)
		felt := true
		for felt {
			ji = (ji + 1) % len(jets)
			rockPush(stall, r, jets[ji])
			felt = rockDown(stall, r)
			//printState(stall, r)
			if !felt {
				stall = updateStall(stall, r)
			}
		}
		// rock cannot go down anymore, update floors
		for _, p := range stall {
			floor[p[0]] = maxInt(floor[p[0]], p[1])
		}
	}
	maxh := 0
	for _, x := range floor {
		maxh = maxInt(maxh, x)
	}
	//printState(stall, nil)
	return maxh
}

func readFile(file string) Jets {
	b, _ := ioutil.ReadFile(file)
	jets := Jets{}
	for _, b := range b {
		if b == '<' {
			jets = append(jets, LEFT)
		} else {
			jets = append(jets, RIGHT)
		}
	}
	return jets
}

func newRock(n int, f Floor) Rock {
	n = n % 5
	h := 0
	for _, x := range f {
		h = maxInt(h, x)
	}
	if n == 0 {
		// ####
		return Rock{
			Point{2, h + 4}, Point{3, h + 4}, Point{4, h + 4}, Point{5, h + 4},
		}
	} else if n == 1 {
		// .#.
		// ###
		// .#.
		return Rock{
			/*            */ Point{3, h + 6},
			Point{2, h + 5}, Point{3, h + 5}, Point{4, h + 5},
			/*            */ Point{3, h + 4},
		}
	} else if n == 2 {
		// ..#
		// ..#
		// ###
		return Rock{
			/*                             */ Point{4, h + 6},
			/*                             */ Point{4, h + 5},
			Point{2, h + 4}, Point{3, h + 4}, Point{4, h + 4},
		}
	} else if n == 3 {
		// #
		// #
		// #
		// #
		return Rock{
			Point{2, h + 7},
			Point{2, h + 6},
			Point{2, h + 5},
			Point{2, h + 4},
		}
	} else if n == 4 {
		// ##
		// ##
		return Rock{
			Point{2, h + 5}, Point{3, h + 5},
			Point{2, h + 4}, Point{3, h + 4},
		}
	}
	return nil
}

func rockPush(stall []Point, r Rock, d int) bool {
	// check first if boundaries prevent the push
	for _, p := range r {
		newx := p[0] + d
		if newx < 0 || newx >= WIDTH {
			return false
		}
		for _, sp := range stall {
			if newx == sp[0] && p[1] == sp[1] {
				return false
			}
		}
	}
	// apply the push
	for i, _ := range r {
		r[i][0] = r[i][0] + d
	}
	return true
}

func rockDown(stall []Point, r Rock) bool {
	// check first if rock can keep falling
	for _, p := range r {
		if p[1] == 1 {
			return false
		}
		for _, sp := range stall {
			if p[0] == sp[0] && p[1]-1 == sp[1] {
				return false
			}
		}
	}
	// apply the fall
	for i, _ := range r {
		r[i][1] = r[i][1] - 1
	}
	return true
}

func updateStall(stall []Point, r Rock) []Point {
	return append(stall, r...)
}

func printState(stall []Point, r Rock) {
	h := 0
	for _, p := range stall {
		h = maxInt(h, p[1])
	}
	for _, p := range r {
		h = maxInt(h, p[1])
	}

	for ; h > 0; h-- {
		fmt.Print("|")
	OUTER:
		for i := 0; i < WIDTH; i++ {
			for _, p := range stall {
				if p[0] == i && p[1] == h {
					fmt.Print("#")
					continue OUTER
				}
			}
			for _, p := range r {
				if p[1] == h && p[0] == i {
					fmt.Print("@")
					continue OUTER
				}
			}
			fmt.Print(".")
		}
		fmt.Print("|\n")
	}
	fmt.Print("+")
	for i := 0; i < WIDTH; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
	fmt.Println()
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

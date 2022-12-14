package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const START_X = 500
const START_Y = 0

type XY struct {
	x int
	y int
}

type State struct {
	rocks      map[XY]bool
	sand       map[XY]bool
	limitLeft  int
	limitRight int
	floor      int
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	state := readFile(file)
	runPart1(state)
	part1 := len(state.sand)

	state.sand = make(map[XY]bool) // reset sand
	runPart2(state)
	part2 := len(state.sand)

	return part1, part2
}

func readFile(file string) *State {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	s := &State{
		rocks:      make(map[XY]bool),
		sand:       make(map[XY]bool),
		limitLeft:  math.MaxInt,
		limitRight: math.MinInt,
		floor:      0,
	}

	for scanner.Scan() {
		addRockPaths(s, scanner.Text())
	}

	readFile.Close()
	return s
}

func addRockPaths(s *State, line string) {
	tracks := strings.Split(line, " -> ")
	l := len(tracks)
	from := parseXY(tracks[0])
	s.addRock(from)
	s.updateLimits(from)
	for i := 1; i < l; i++ {
		to := parseXY(tracks[i])
		dx := getDelta(from.x, to.x)
		dy := getDelta(from.y, to.y)
		for {
			from.x += dx
			from.y += dy
			s.addRock(from)
			s.updateLimits(from)
			if from.x == to.x && from.y == to.y {
				break
			}
		}
	}
}

func parseXY(str string) *XY {
	coords := strings.Split(str, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return &XY{x, y}
}

func getDelta(from int, to int) int {
	if from < to {
		return 1
	} else if from > to {
		return -1
	} else {
		return 0
	}
}

func (s *State) updateLimits(rock *XY) {
	if rock.x < s.limitLeft {
		s.limitLeft = rock.x
	}
	if rock.x > s.limitRight {
		s.limitRight = rock.x
	}
	if rock.y+2 > s.floor {
		s.floor = rock.y + 2
	}
}

func (s *State) addRock(rock *XY) {
	key := *rock // create a copy, rock may be modified outside
	s.rocks[key] = true
}

func (s *State) addSand(sand *XY) {
	// we know this sand will not be modified and can avoid the copy
	s.sand[*sand] = true
}

func runPart1(s *State) {
	for {
		sand := &XY{START_X, START_Y}
		for {
			if sand.x < s.limitLeft || sand.y > s.limitRight {
				// sand will fall forever, we are done
				return
			}
			if s.isEmpty(sand.x, sand.y+1, true) {
				// move down
				sand.y++
			} else if s.isEmpty(sand.x-1, sand.y+1, true) {
				// move left-down
				sand.x--
				sand.y++
			} else if s.isEmpty(sand.x+1, sand.y+1, true) {
				// move right-down
				sand.x++
				sand.y++
			} else {
				// time to rest
				s.addSand(sand)
				break
			}
		}
	}
}

func runPart2(s *State) {
	for {
		sand := &XY{START_X, START_Y}
		if !s.isEmpty(sand.x, sand.y, false) {
			// sand source is blocked, we are done
			return
		}
		for {
			if s.isEmpty(sand.x, sand.y+1, false) {
				// move down
				sand.y++
			} else if s.isEmpty(sand.x-1, sand.y+1, false) {
				// move left-down
				sand.x--
				sand.y++
			} else if s.isEmpty(sand.x+1, sand.y+1, false) {
				// move right-down
				sand.x++
				sand.y++
			} else {
				// time to rest
				s.addSand(sand)
				break
			}
		}
	}
}

func (s *State) isEmpty(x int, y int, ignoreFloor bool) bool {
	xy := XY{x, y}
	return !s.rocks[xy] && !s.sand[xy] && (ignoreFloor || s.floor > y)
}

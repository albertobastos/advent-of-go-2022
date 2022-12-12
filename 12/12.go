package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const INFINITY = math.MaxInt

type State struct {
	heights   []int
	distances []int
	visited   []bool
	width     int
	start     int
	end       int
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	s := readFile(file)
	s.fillDistances()
	part1 := s.distances[s.end]

	s.reset()
	lowest := int('a') - '0'
	for i, h := range s.heights {
		if h == lowest {
			s.distances[i] = 0
		}
	}
	s.fillDistances()
	part2 := s.distances[s.end]

	return part1, part2
}

func readFile(file string) *State {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	s := State{[]int{}, []int{}, []bool{}, -1, -1, -1}

	for scanner.Scan() {
		line := scanner.Text()
		if s.width == -1 {
			s.width = len(line)
		}
		for _, c := range line {
			s.heights = append(s.heights, runeToHeight(c))
			s.distances = append(s.distances, INFINITY)
			s.visited = append(s.visited, false)
			if c == 'S' {
				s.start = len(s.distances) - 1
				s.distances[s.start] = 0
			} else if c == 'E' {
				s.end = len(s.distances) - 1
			}
		}
	}

	readFile.Close()
	return &s
}

func runeToHeight(c rune) int {
	if c == 'S' {
		c = 'a'
	} else if c == 'E' {
		c = 'z'
	}
	return int(c) - '0'
}

func getNextCandidate(s *State) int {
	min_i := -1
	min_d := INFINITY

	for i, d := range s.distances {
		if !s.visited[i] {
			if min_i == -1 || min_d > s.distances[i] {
				min_i, min_d = i, d
			}
		}
	}

	return min_i
}

func getMovementsFrom(s *State, pos int) []int {
	moves := []int{}
	currh := s.heights[pos]

	// left?
	if pos%s.width > 0 && s.heights[pos-1] <= currh+1 && !s.visited[pos-1] {
		moves = append(moves, pos-1)
	}

	// right?
	if pos%s.width < s.width-1 && s.heights[pos+1] <= currh+1 && !s.visited[pos+1] {
		moves = append(moves, pos+1)
	}

	// up?
	if pos-s.width >= 0 && s.heights[pos-s.width] <= currh+1 && !s.visited[pos-s.width] {
		moves = append(moves, pos-s.width)
	}

	// down?
	if pos+s.width < len(s.heights) && s.heights[pos+s.width] <= currh+1 && !s.visited[pos+s.width] {
		moves = append(moves, pos+s.width)
	}

	return moves
}

func (s *State) print(title string, fn func(*State, int) string) {
	fmt.Println(title)
	for i, _ := range s.heights {
		fmt.Printf("%s ", fn(s, i))
		if (i+1)%s.width == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (s *State) fillDistances() {
	for {
		curr := getNextCandidate(s)
		if curr == -1 {
			break
		}
		newd := s.distances[curr] + 1
		for _, move := range getMovementsFrom(s, curr) {
			if s.distances[move] > newd {
				s.distances[move] = newd
			}
		}
		s.visited[curr] = true
		//s.printVisited()
		//s.printDistances()
	}
}

func (s *State) reset() {
	for i, _ := range s.heights {
		s.distances[i] = INFINITY
		s.visited[i] = false
	}
	s.distances[s.start] = 0
}

func (s *State) printDistances() {
	s.print("Distances", func(t *State, pos int) string {
		d := s.distances[pos]
		if d == INFINITY {
			return "**"
		} else {
			return fmt.Sprintf("%02d", d)
		}
	})
}

func (s *State) printVisited() {
	s.print("Visited", func(t *State, pos int) string {
		if s.visited[pos] {
			return "X"
		} else {
			return "."
		}
	})
}

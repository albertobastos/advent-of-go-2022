package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const RIGHT = "R"
const UP = "U"
const DOWN = "D"
const LEFT = "L"

type XY struct {
	x int
	y int
}

type State struct {
	knots        []XY // [head..tail]
	visited_tail map[string]bool
}

type Move struct {
	dir string
	n   int
}

func main() {
	fmt.Println("Part1:", run("input.txt", 2))
	fmt.Println("Part2:", run("input.txt", 10))
}

func run(file string, knots int) int {
	moves := readFile(file)
	state := initState(knots)
	state.runMoves(moves)
	return len(state.visited_tail)
}

func readFile(file string) []Move {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, " ")
		n, _ := strconv.Atoi(spl[1])
		moves = append(moves, Move{spl[0], n})
	}

	readFile.Close()
	return moves
}

func initState(knots int) *State {
	s := State{
		knots:        initKnots(knots),
		visited_tail: make(map[string]bool),
	}
	s.updateVisited()
	return &s
}

func initKnots(knots int) []XY {
	k := []XY{}
	for i := 0; i < knots; i++ {
		k = append(k, XY{0, 0})
	}
	return k
}

func (s *State) updateVisited() {
	s.visited_tail[s.knots[len(s.knots)-1].toKey()] = true
}

func (xy *XY) toKey() string {
	return fmt.Sprintf("%d,%d", xy.x, xy.y)
}

func (s *State) runMoves(moves []Move) {
	head := s.getHead()
	for _, m := range moves {
		dx, dy := m.getDeltas()
		for i := 0; i < m.n; i++ {
			head.x += dx
			head.y += dy
			for j := 0; j < len(s.knots)-1; j++ {
				s.maybeMove(s.getKnotAt(j), s.getKnotAt(j+1))
			}
			s.updateVisited()
		}
		//s.print()
	}
}

func (s *State) getHead() *XY {
	return s.getKnotAt(0)
}

func (s *State) getKnotAt(i int) *XY {
	return &s.knots[i]
}
func (m Move) getDeltas() (int, int) {
	if m.dir == RIGHT {
		return 1, 0
	} else if m.dir == UP {
		return 0, 1
	} else if m.dir == LEFT {
		return -1, 0
	} else if m.dir == DOWN {
		return 0, -1
	}
	fmt.Println("Invalid move:", m)
	os.Exit(1)
	return 0, 0
}

func (s *State) maybeMove(head *XY, tail *XY) bool {
	nx := tail.x + getDelta(head.x, tail.x)
	ny := tail.y + getDelta(head.y, tail.y)
	if head.x != nx || head.y != ny {
		tail.x, tail.y = nx, ny
		return true
	} else if head.x != nx {
		tail.x = nx
		return true
	} else if head.y != ny {
		tail.y = ny
		return true
	} else {
		return false
	}
}

func getDelta(from int, to int) int {
	dist := from - to
	if dist < 0 {
		return -1
	} else if dist > 0 {
		return 1
	} else {
		return 0
	}
}

func (s *State) print() {
	fmt.Println("State:")
	for i, k := range s.knots {
		label := fmt.Sprint(i)
		if i == 0 {
			label = "H"
		}
		fmt.Println("knot", label, "[", k.x, k.y, "]")
	}
	fmt.Println()
}

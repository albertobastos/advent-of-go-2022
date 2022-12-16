package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const INIT_MINUTES = 30
const INIT_VALVE = ValveName("AA")

var RE_VALVE = regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ([A-Z, ]+)")

type ValveName string
type Valve struct {
	name      ValveName
	flowRate  int
	opened    bool
	distances map[ValveName]int
}

type State struct {
	valves      map[ValveName]*Valve
	minutesLeft int
	current     ValveName
}

type ValveList []ValveName

func (l ValveList) indexOf(name ValveName) int {
	for i, n := range l {
		if n == name {
			return i
		}
	}
	return -1
}

func (l ValveList) append(name ValveName) ValveList {
	return append(l, name)
}

func (l ValveList) remove(name ValveName) ValveList {
	toRemove := l.indexOf(name)
	if toRemove < 0 {
		fmt.Println("Error: trying to remove Valve not in ValveList")
		os.Exit(1)
	}

	cpy := ValveList{}
	cpy = append(cpy, l[:toRemove]...)
	cpy = append(cpy, l[toRemove+1:]...)
	return cpy
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	state := readFile(file)

	unbroken := filterNotBroken(state)
	_, released := findOptimalRelease(
		state, unbroken, ValveList{INIT_VALVE}, INIT_MINUTES, 0)

	//fmt.Println("Visited:", visited) // AA, DD, BB, JJ, HH, EE, CC

	return released, -1
}

func readFile(file string) *State {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	s := &State{
		valves:      make(map[ValveName]*Valve),
		minutesLeft: INIT_MINUTES,
		current:     INIT_VALVE,
	}
	for scanner.Scan() {
		readValve(s, scanner.Text())
	}

	fillDistances(s)

	readFile.Close()
	return s
}

func readValve(s *State, str string) {
	m := RE_VALVE.FindStringSubmatch(str)
	name := ValveName(m[1])
	flowRate, _ := strconv.Atoi(m[2])
	conns := strings.Split(m[3], ", ")
	v := Valve{
		name:      name,
		flowRate:  flowRate,
		opened:    false,
		distances: make(map[ValveName]int)}
	v.distances[v.name] = 0
	for _, conn := range conns {
		v.distances[ValveName(conn)] = 1
	}
	s.valves[v.name] = &v
}

func fillDistances(s *State) {
	// init all unknown distances with the max possible value
	nvalves := len(s.valves)
	for _, v := range s.valves {
		for name, _ := range s.valves {
			if v.name != name && v.distances[name] == 0 {
				v.distances[name] = nvalves
			}
		}
	}

	// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
	for k, vk := range s.valves {
		for _, vi := range s.valves {
			for j, _ := range s.valves {
				vi.distances[j] = minInt(vi.distances[j], vi.distances[k]+vk.distances[j])
			}
		}
	}
}

func filterNotBroken(s *State) ValveList {
	l := ValveList{}
	for _, v := range s.valves {
		if v.flowRate > 0 {
			l = append(l, v.name)
		}
	}
	return l
}

func findOptimalRelease(s *State, closed ValveList, visited ValveList, minutesLeft int, releasedSoFar int) (ValveList, int) {
	if len(closed) == 0 || minutesLeft == 0 {
		return visited, releasedSoFar
	}
	current := s.valves[visited[len(visited)-1]]
	bestVisited := visited
	bestReleased := releasedSoFar
	for _, vname := range closed {
		minutesLeftAfterOpen := minutesLeft - current.distances[vname] - 1

		if minutesLeftAfterOpen >= 0 {
			iVisited, iReleased := findOptimalRelease(
				s,
				closed.remove(vname),
				visited.append(vname),
				minutesLeftAfterOpen,
				releasedSoFar+minutesLeftAfterOpen*s.valves[vname].flowRate,
			)
			if iReleased > bestReleased {
				bestVisited = iVisited
				bestReleased = iReleased
			}
		}
	}

	return bestVisited, bestReleased
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

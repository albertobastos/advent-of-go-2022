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
	distances map[ValveName]int
}

type State struct {
	valves  map[ValveName]*Valve
	current ValveName
}

type ValveSet map[ValveName]bool

func vadd(s ValveSet, name ValveName) ValveSet {
	if s[name] {
		return s
	}
	cpy := make(ValveSet)
	for n, _ := range s {
		cpy[n] = true
	}
	cpy[name] = true
	return cpy
}

func vdel(s ValveSet, name ValveName) ValveSet {
	if !s[name] {
		return s
	}
	cpy := make(ValveSet)
	for n, _ := range s {
		if n != name {
			cpy[n] = true
		}
	}
	return cpy
}

func vinit(names ...ValveName) ValveSet {
	s := make(ValveSet)
	for _, n := range names {
		s[n] = true
	}
	return s
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	state := readFile(file)

	unbroken := state.getValvesWithFlow()
	released := findOptimalRelease(
		state, unbroken, vinit(INIT_VALVE), state.valves[INIT_VALVE], INIT_MINUTES, 0)
	part2 := doPart2(state, unbroken)

	return released, part2
}

func readFile(file string) *State {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	s := &State{
		valves:  make(map[ValveName]*Valve),
		current: INIT_VALVE,
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

func (s *State) getValvesWithFlow() ValveSet {
	r := vinit()
	for _, v := range s.valves {
		if v.flowRate > 0 {
			r[v.name] = true
		}
	}
	return r
}

func findOptimalRelease(s *State, closed ValveSet, visited ValveSet, current *Valve, minutesLeft int, releasedSoFar int) int {
	if len(closed) == 0 || minutesLeft == 0 {
		return releasedSoFar
	}
	bestReleased := releasedSoFar
	for vname, _ := range closed {
		minutesLeftAfterOpen := minutesLeft - current.distances[vname] - 1

		if minutesLeftAfterOpen >= 0 {
			v := s.valves[vname]
			iReleased := findOptimalRelease(
				s,
				vdel(closed, vname),
				vadd(visited, vname),
				v,
				minutesLeftAfterOpen,
				releasedSoFar+minutesLeftAfterOpen*v.flowRate,
			)
			if iReleased > bestReleased {
				bestReleased = iReleased
			}
		}
	}

	return bestReleased
}

func doPart2(s *State, flowable ValveSet) int {
	return 0
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

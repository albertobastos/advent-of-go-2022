package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const UNKNOWN = 0
const EMPTY = 1
const BEACON = 2

var (
	RE_SENSOR = regexp.MustCompile("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)")
)

type Track [2]int
type Tracks map[int][]Track       // row -> empty tracks
type Beacons map[int]map[int]bool // [row,col] -> has beacon
type State struct {
	tracks  Tracks
	beacons Beacons
}

func main() {
	part1, part2 := run(2000000, 4000000, "input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(row int, distressmax int, file string) (int, int) {
	s := processFile(file)
	part1 := doPart1(s, row)
	part2 := doPart2(s, distressmax)
	return part1, part2
}

func processFile(file string) *State {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	s := &State{make(Tracks), make(Beacons)}
	for scanner.Scan() {
		processSensor(s, scanner.Text())
	}

	readFile.Close()
	return s
}

func processSensor(s *State, log string) {
	mxs := RE_SENSOR.FindStringSubmatch(log)
	sx, _ := strconv.Atoi(mxs[1])
	sy, _ := strconv.Atoi(mxs[2])
	bx, _ := strconv.Atoi(mxs[3])
	by, _ := strconv.Atoi(mxs[4])
	d := getDistance(sx, sy, bx, by)

	s.addBeacon(bx, by)
	for y := sy - d; y <= sy+d; y++ {
		dx := d - abs(sy-y)
		from := sx - dx
		to := sx + dx
		if y == by {
			if from == bx {
				from++
			}
			if to == bx {
				to--
			}
		}
		if to-from >= 0 {
			s.addTrack(y, from, to)
		}
	}
}

func (s *State) addBeacon(x int, y int) {
	if s.beacons[y] == nil {
		s.beacons[y] = make(map[int]bool)
	}
	s.beacons[y][x] = true
}

func (s *State) hasBeacon(x int, y int) bool {
	if s.beacons[y] == nil {
		return false
	}
	return s.beacons[y][x]
}

func (s *State) addTrack(row int, from int, to int) {
	rt := s.tracks[row]
	if s.tracks[row] == nil {
		rt = []Track{}
	}

	rt = append(rt, [2]int{from, to})
	s.tracks[row] = rt
}

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func doPart1(s *State, row int) int {
	rts := s.tracks[row]
	if rts == nil {
		return 0
	}

	sortTracks(rts)

	sum := 0
	prev := math.MinInt
	for _, t := range rts {
		from := maxInt(prev+1, t[0])
		to := t[1]
		if to >= from {
			sum += to - from + 1
			prev = to
		}
	}
	return sum
}

func doPart2(s *State, distressmax int) int {
	bx, by := findDistressBeacon(s, distressmax)
	return bx*distressmax + by
}

func findDistressBeacon(s *State, distressmax int) (int, int) {
	for y := 0; y <= distressmax; y++ {
		yts := s.tracks[y]
		sortTracks(yts)
		last := -1
		for _, t := range yts {
			for x := last + 1; x < t[0] && x <= distressmax; x++ {
				if !s.hasBeacon(x, y) {
					return x, y
				}
			}
			last = maxInt(last, t[1])
			if last > distressmax {
				// remaining tracks are past the max allowed, skip row
				break
			}
		}
	}
	return 0, 0
}

func sortTracks(ts []Track) {
	sort.SliceStable(ts, func(i, j int) bool {
		return ts[i][0] <= ts[j][0]
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const FREQ_MUL = 4000000

type Sensor struct {
	x  int // sensor x
	y  int // sensor y
	bx int // beacon x
	by int // beacon y
	d  int // distance from sensor to beacon
}
type Track [2]int
type TrackList []Track

var RE_SENSOR = regexp.MustCompile("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)")
var NO_TRACK = Track{math.MaxInt, math.MaxInt}

func main() {
	part1, part2 := run(2000000, 4000000, "input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(row int, distressmax int, file string) (int, int) {
	sensors := readFile(file)
	part1 := doPart1(sensors, row)
	part2 := doPart2(sensors, distressmax)
	return part1, part2
}

func readFile(file string) []*Sensor {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	list := []*Sensor{}
	for scanner.Scan() {
		list = append(list, readSensor(scanner.Text()))
	}

	readFile.Close()
	return list
}

func readSensor(log string) *Sensor {
	mxs := RE_SENSOR.FindStringSubmatch(log)
	x, _ := strconv.Atoi(mxs[1])
	y, _ := strconv.Atoi(mxs[2])
	bx, _ := strconv.Atoi(mxs[3])
	by, _ := strconv.Atoi(mxs[4])
	return &Sensor{x, y, bx, by, distance(x, y, bx, by)}
}

func doPart1(sensors []*Sensor, row int) int {
	tracks := TrackList{}
	for _, s := range sensors {
		t := s.getTrack(row, false)
		if t != NO_TRACK {
			tracks = mergeTracks(tracks, t)
		}
	}

	sum := 0
	for _, t := range tracks {
		sum += t[1] - t[0] + 1
	}
	return sum
}

func doPart2(sensors []*Sensor, distressmax int) int {
	x, y := findDistressBeacon(sensors, distressmax)
	return x*FREQ_MUL + y
}

func findDistressBeacon(sensors []*Sensor, distressmax int) (int, int) {
	for row := 0; row <= distressmax; row++ {
		tracks := TrackList{}
		for _, s := range sensors {
			t := s.getTrack(row, true)
			if t != NO_TRACK {
				tracks = mergeTracks(tracks, t)
			}
		}
		ltracks := len(tracks)
		if ltracks == 0 {
			fmt.Println("Easy case 1, suspicious...")
			return 0, row
		}
		if tracks[0][0] > 0 {
			fmt.Println("Easy case 2, suspicious...")
			return 0, row
		}
		if tracks[ltracks-1][1] < distressmax {
			fmt.Println("Easy case 3, suspicious...")
			return distressmax, row
		}
		for i := 0; i < ltracks-1; i++ {
			cand := tracks[i][1] + 1
			if cand < tracks[i+1][0] {
				// found it!
				return cand, row
			}
		}
	}
	return 0, 0
}

func mergeTracks(tracks TrackList, track Track) TrackList {
	// insert it ordered by track start
	i := 0
	for i < len(tracks) && tracks[i][0] < track[0] {
		i++
	}
	tracks = insert(tracks, i, track)

	// merge overlapping tracks
	ntracks := TrackList{tracks[0]}
	for _, t := range tracks[1:] {
		curr := ntracks[len(ntracks)-1]
		if t[0] > curr[1] {
			// no overlap, append it
			ntracks = append(ntracks, t)
		} else {
			// overlap, extend to the right
			curr[1] = maxInt(t[1], curr[1])
			ntracks[len(ntracks)-1] = curr
		}
	}

	return ntracks
}

func insert(arr TrackList, index int, elem Track) TrackList {
	if len(arr) == index {
		return append(arr, elem)
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = elem
	return arr
}

func (s *Sensor) getTrack(row int, withBeacon bool) Track {
	if s.y-s.d > row || s.y+s.d < row {
		return NO_TRACK
	}
	if !withBeacon && row == s.by && abs(s.y-row) == s.d {
		// the beacon is the whole track
		return NO_TRACK
	}
	d := s.d - abs(s.y-row)
	from, to := s.x-d, s.x+d
	if !withBeacon && s.by == row {
		if s.bx == from {
			from++
		} else if s.bx == to {
			to--
		}
	}
	return Track{from, to}
}

// auxiliar functions

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
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

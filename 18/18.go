package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube [3]int // x, y, z
func (c Cube) siblings() [6]Cube {
	return [6]Cube{
		{c[0] - 1, c[1], c[2]},
		{c[0] + 1, c[1], c[2]},
		{c[0], c[1] - 1, c[2]},
		{c[0], c[1] + 1, c[2]},
		{c[0], c[1], c[2] - 1},
		{c[0], c[1], c[2] + 1},
	}
}

type CubeSet map[Cube]bool

func (s CubeSet) add(c Cube) {
	s[c] = true
}
func (s CubeSet) has(c Cube) bool {
	return s[c]
}
func (s CubeSet) del(c Cube) bool {
	if s[c] {
		delete(s, c)
		return true
	}
	return false
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	cubes := readFile(file)

	freeSides := 0
	visibleSides := 0
	for c, _ := range cubes {
		for _, s := range c.siblings() {
			if !cubes.has(s) {
				freeSides++
			}
		}
		//visibleSides += findVisibleSides(cubes, c)
	}

	return freeSides, visibleSides
}

func readFile(file string) CubeSet {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	cubes := make(CubeSet)

	for scanner.Scan() {
		cubes.add(readCube(scanner.Text()))
	}

	readFile.Close()
	return cubes
}

func readCube(str string) Cube {
	spl := strings.Split(str, ",")
	x, _ := strconv.Atoi(spl[0])
	y, _ := strconv.Atoi(spl[1])
	z, _ := strconv.Atoi(spl[2])
	return Cube{x, y, z}
}

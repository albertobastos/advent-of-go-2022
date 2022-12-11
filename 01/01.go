package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	elves := readFile(file)
	sort.Ints(elves)

	part1 := elves[len(elves)-1]

	sum := 0
	for _, calories := range elves[len(elves)-3:] {
		sum += calories
	}
	part2 := sum

	return part1, part2
}

func readFile(file string) []int {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	accs := []int{0}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			accs = append(accs, 0)
		} else {
			lineInt, _ := strconv.Atoi(line)
			accs[len(accs)-1] += lineInt
		}
	}

	readFile.Close()
	return accs
}

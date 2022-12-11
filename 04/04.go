package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coverage [2]int
type Pair [2]Coverage

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	pairs := readFile(file)
	fullOverlaps := 0
	someOverlaps := 0
	for _, pair := range pairs {
		fullOverlaps += pair.hasFullOverlap()
		someOverlaps += pair.hasSomeOverlap()
	}
	return fullOverlaps, someOverlaps
}

func readFile(file string) []Pair {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	pairs := []Pair{}

	for scanner.Scan() {
		line := scanner.Text()
		covs_str := strings.Split(line, ",")
		if len(covs_str) != 2 {
			fmt.Println("Invalid coverage pair:", line)
			os.Exit(1)
		}
		pair := Pair{}
		for cov_i, cov_str := range covs_str {
			pair_str := strings.Split(cov_str, "-")
			if len(pair_str) != 2 {
				fmt.Println("Invalid coverage:", cov_str)
				os.Exit(1)
			}
			cov := Coverage{}
			cov[0], _ = strconv.Atoi(pair_str[0])
			cov[1], _ = strconv.Atoi(pair_str[1])
			pair[cov_i] = cov
		}
		pairs = append(pairs, pair)
	}

	readFile.Close()
	return pairs
}

func (p Pair) hasFullOverlap() int {
	if p[0].from() >= p[1].from() && p[0].to() <= p[1].to() {
		return 1
	} else if p[1].from() >= p[0].from() && p[1].to() <= p[0].to() {
		return 1
	} else {
		return 0
	}
}

func (p Pair) hasSomeOverlap() int {
	if p[0].from() <= p[1].to() && p[0].to() >= p[1].from() {
		return 1
	} else {
		return 0
	}
}

func (c Coverage) from() int {
	return c[0]
}

func (c Coverage) to() int {
	return c[1]
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	elves := readFile()
	sort.Ints(elves)

	fmt.Println("Part1:", elves[len(elves)-1], "calories.")

	sum := 0
	for _, calories := range elves[len(elves)-3:] {
		sum += calories
	}
	fmt.Println("Part2:", sum, "calories.")
}

func readFile() []int {
	readFile, _ := os.Open("input.txt")
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

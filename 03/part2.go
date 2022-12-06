package main

import (
	"bufio"
	"fmt"
	"os"
)

const GROUP_SIZE = 3

type Elf string
type BadgeMap map[rune]bool

func doPart2(file string) int {
	elves := p2_readFile(file)
	sum := 0
	for i := 0; i < len(elves); i = i + GROUP_SIZE {
		sum += int(priority(findBadge(elves[i : i+GROUP_SIZE])))
	}
	return sum
}

func p2_readFile(file string) []Elf {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	list := []Elf{}

	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, Elf(line))
	}

	readFile.Close()
	return list
}

func findBadge(elves []Elf) rune {
	badges := [](BadgeMap){}
	for _, e := range elves {
		badges = append(badges, badgesMap(e))
	}
	for curr := range badges[0] {
		all := true
		for _, m := range badges[1:] {
			_, ok := m[curr]
			all = all && ok
		}
		if all {
			return curr
		}
	}
	fmt.Println("No common badge found for elves:", elves)
	return -1
}

func badgesMap(e Elf) BadgeMap {
	m := make(BadgeMap)
	for _, c := range e {
		m[c] = true
	}
	return m
}

package main

import (
	"bufio"
	"os"
)

type Pocket map[Priority]bool
type Rucksack struct {
	pocket1 *Pocket
	pocket2 *Pocket
}

func doPart1(file string) int {
	rucksacks := p1_readFile(file)
	sum := 0
	for _, r := range rucksacks {
		for _, p := range r.findCommon() {
			sum += int(p)
		}
	}
	return sum
}

func p1_readFile(file string) []Rucksack {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	list := []Rucksack{}

	for scanner.Scan() {
		line := scanner.Text()
		r := Rucksack{}
		r.parse(line)
		list = append(list, r)
	}

	readFile.Close()
	return list
}

func (r *Rucksack) parse(str string) {
	pocketSize := len(str) / 2
	pocket1 := make(Pocket)
	pocket2 := make(Pocket)
	pocket1.parse(str[:pocketSize])
	pocket2.parse(str[pocketSize:])
	r.pocket1 = &pocket1
	r.pocket2 = &pocket2
}

func (p Pocket) parse(str string) {
	for _, c := range str {
		p[priority(c)] = true
	}
}

func (r *Rucksack) findCommon() []Priority {
	common := []Priority{}
	for p1 := range *(r.pocket1) {
		for p2 := range *(r.pocket2) {
			if p1 == p2 {
				common = append(common, p1)
			}
		}
	}
	return common
}

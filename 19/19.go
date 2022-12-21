package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var RE_BLUEPRINT = regexp.MustCompile("Blueprint ([0-9]+): Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.")

type Blueprint struct {
	id                   int
	robotOre_ores        int
	robotClay_ores       int
	robotObsidian_ores   int
	robotObsidian_clays  int
	robotGeode_ores      int
	robotGeode_obsidians int
}

type Inventory struct {
	ore            int
	clay           int
	obsidian       int
	geode          int
	robotsOre      int
	robotsClay     int
	robotsObsidian int
	robotsGeode    int
}

func (inv *Inventory) clone() *Inventory {
	cpy := Inventory{
		inv.ore,
		inv.clay,
		inv.obsidian,
		inv.geode,
		inv.robotsOre,
		inv.robotsClay,
		inv.robotsObsidian,
		inv.robotsGeode,
	}
	return &cpy
}

func main() {
	part1, part2 := run("input.txt", 24)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string, minutes int) (int, int) {
	blueprints := readFile(file)
	qualityLevelSum := 0
	for _, b := range blueprints {
		geodes := maximumGeodes(b, minutes)
		fmt.Println("bp", b.id, "geodes", geodes)
		qualityLevelSum += (b.id * geodes)
	}
	return qualityLevelSum, -1
}

func readFile(file string) []*Blueprint {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	list := []*Blueprint{}

	for scanner.Scan() {
		list = append(list, readBlueprint(scanner.Text()))
	}

	readFile.Close()
	return list
}

func readBlueprint(str string) *Blueprint {
	m := RE_BLUEPRINT.FindStringSubmatch(str)
	mi := allToInt(m[1:])
	return &Blueprint{
		id:                   mi[0],
		robotOre_ores:        mi[1],
		robotClay_ores:       mi[2],
		robotObsidian_ores:   mi[3],
		robotObsidian_clays:  mi[4],
		robotGeode_ores:      mi[5],
		robotGeode_obsidians: mi[6],
	}
}

func allToInt(l []string) []int {
	li := []int{}
	for _, s := range l {
		i, _ := strconv.Atoi(s)
		li = append(li, i)
	}
	return li
}

func maximumGeodes(b *Blueprint, minutes int) int {
	inv := Inventory{
		robotsOre: 1,
	}
	return maximumGeodesDPS(b, &inv, minutes, 0)
}

func maximumGeodesDPS(b *Blueprint, inv *Inventory, minutesLeft int) int {
	fmt.Println("bp", b.id, ",", minutesLeft, "minutes left", inv)
	if minutesLeft == 0 {
		return inv.geode
	}

	max := 0

	if canBuildOreRobot(b, inv) {
		next := inv.clone()
		collect(next)
		buildOreRobot(b, next)
		max = maxInt(max, maximumGeodesDPS(b, next, minutesLeft-1))
	}

	if canBuildClayRobot(b, inv) {
		next := inv.clone()
		collect(next)
		buildClayRobot(b, next)
		max = maxInt(max, maximumGeodesDPS(b, next, minutesLeft-1))
	}

	if canBuildObsidianRobot(b, inv) {
		next := inv.clone()
		collect(next)
		buildObsidianRobot(b, next)
		max = maxInt(max, maximumGeodesDPS(b, next, minutesLeft-1))
	}

	if canBuildGeodeRobot(b, inv) {
		next := inv.clone()
		collect(next)
		buildGeodeRobot(b, next)
		max = maxInt(max, maximumGeodesDPS(b, next, minutesLeft-1))
	}

	// no building during this minute
	next := inv.clone()
	collect(next)
	max = maxInt(max, maximumGeodesDPS(b, next, minutesLeft-1))

	return max
}

func collect(inv *Inventory) {
	inv.ore += inv.robotsOre
	inv.clay += inv.robotsClay
	inv.obsidian += inv.robotsObsidian
	inv.geode += inv.robotsGeode
}

func canBuildOreRobot(b *Blueprint, inv *Inventory) bool {
	return inv.ore >= b.robotOre_ores
}

func canBuildClayRobot(b *Blueprint, inv *Inventory) bool {
	return inv.ore >= b.robotClay_ores
}

func canBuildObsidianRobot(b *Blueprint, inv *Inventory) bool {
	return inv.ore >= b.robotObsidian_ores && inv.clay >= b.robotObsidian_clays
}

func canBuildGeodeRobot(b *Blueprint, inv *Inventory) bool {
	return inv.ore >= b.robotGeode_ores && inv.obsidian >= b.robotGeode_obsidians
}

func buildOreRobot(b *Blueprint, inv *Inventory) {
	inv.ore -= b.robotOre_ores
	inv.robotsOre++
}

func buildClayRobot(b *Blueprint, inv *Inventory) {
	inv.ore -= b.robotClay_ores
	inv.robotsClay++
}

func buildObsidianRobot(b *Blueprint, inv *Inventory) {
	inv.ore -= b.robotObsidian_ores
	inv.clay -= b.robotObsidian_clays
	inv.robotsObsidian++
}

func buildGeodeRobot(b *Blueprint, inv *Inventory) {
	inv.ore -= b.robotGeode_ores
	inv.obsidian -= b.robotGeode_obsidians
	inv.robotsGeode++
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

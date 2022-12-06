package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(file string) string {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//line := scanner.Text()
		// TODO
	}

	readFile.Close()
	return "TODO"
}

func run(file string) (int, int) {
	input := readFile(file)
	// TODO
	fmt.Println(input)
	return -1, -1
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

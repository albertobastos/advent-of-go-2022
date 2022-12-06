package main

import (
	"fmt"
	"os"
)

const PMARKER_LEN = 4
const MMARKER_LEN = 14

func readFile() string {
	data, _ := os.ReadFile("input.txt")
	return string(data)
}

func _findStart(str string, ml int) int {
	i := 0
	l := len(str)
	for i+ml < l {
		if allUnique(str[i : i+ml]) {
			return i
		}
		i++
	}
	return -1
}

func findPacketStart(str string) int {
	return _findStart(str, PMARKER_LEN)
}

func findMessageStart(str string) int {
	return _findStart(str, MMARKER_LEN)
}

func allUnique(str string) bool {
	set := make(map[rune]bool)
	for _, c := range str {
		if set[c] {
			return false
		}
		set[c] = true
	}
	return true
}

func main() {
	input := readFile()
	fmt.Println("Part1:", findPacketStart(input)+PMARKER_LEN)
	fmt.Println("Part2:", findMessageStart(input)+MMARKER_LEN)
}

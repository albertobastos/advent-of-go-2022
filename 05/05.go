package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack []rune
type Move string

var RE_NUM *regexp.Regexp

func readFile() ([]Stack, []Move) {
	readFile, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	var stacks []Stack
	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 4 && line[0:4] == "move" {
			// move line
			moves = append(moves, Move(line))
		} else if len(line) == 0 {
			// empty line separator, skip
		} else if len(line) > 1 && line[1] == '1' {
			// stack index line, skip
		} else {
			// stack line
			if stacks == nil {
				stacks = initStacks(line)
			}
			for i, stack := range stacks {
				pos := i*4 + 1
				container := line[pos]
				if container != ' ' {
					stacks[i] = append(stack, rune(container))
				}
			}
		}
	}

	for _, stack := range stacks {
		stack.reverse()
	}

	return stacks, moves
}

func initStacks(s string) []Stack {
	// determines how many stacks do we need based on stack-line length
	stacks := []Stack{}
	howMany := (len(s) / 4) + 1
	for i := 0; i < howMany; i++ {
		stacks = append(stacks, Stack{})
	}
	return stacks
}

func (s Stack) reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func applyMove_mover9000(stacks []Stack, move Move) {
	nums := RE_NUM.FindAllString(string(move), -1)
	count, _ := strconv.Atoi(nums[0])
	from, _ := strconv.Atoi(nums[1])
	to, _ := strconv.Atoi(nums[2])

	// 0-based indexes...
	from--
	to--

	from_len := len(stacks[from])
	for i := 1; i <= count; i++ {
		stacks[to] = append(stacks[to], stacks[from][from_len-i])
	}
	stacks[from] = stacks[from][:from_len-count]
}

func applyMove_mover9001(stacks []Stack, move Move) {
	nums := RE_NUM.FindAllString(string(move), -1)
	count, _ := strconv.Atoi(nums[0])
	from, _ := strconv.Atoi(nums[1])
	to, _ := strconv.Atoi(nums[2])

	// 0-based indexes...
	from--
	to--

	from_len := len(stacks[from])
	stacks[to] = append(stacks[to], stacks[from][from_len-count:]...)
	stacks[from] = stacks[from][:from_len-count]
}

func formatTops(stacks []Stack) string {
	tops := ""
	for _, stack := range stacks {
		tops += string(stack[len(stack)-1])
	}
	return tops
}

func main() {
	RE_NUM = regexp.MustCompile("[0-9]+")
	stacks_part1, moves := readFile()
	stacks_part2, _ := readFile()
	for _, move := range moves {
		applyMove_mover9000(stacks_part1, move)
		applyMove_mover9001(stacks_part2, move)
	}
	fmt.Println("Part1:", formatTops(stacks_part1))
	fmt.Println("Part2:", formatTops(stacks_part2))
}

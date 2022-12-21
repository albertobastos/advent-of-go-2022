package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var RE_VAL = regexp.MustCompile("([a-z]{4}): ([0-9]+)")
var RE_PEND = regexp.MustCompile("([a-z]{4}): ([a-z]{4}) (.) ([a-z]{4})")

const SUM = '+'
const SUB = '-'
const MUL = '*'
const DIV = '/'

type Monkey struct {
	id       string
	op1      string
	op2      string
	operator rune
}

type Pending map[string]*Monkey
type Values map[string]int

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	vals, pend := readFile(file)

	_, ok := vals["root"]

	if !ok {
	OUTER:
		for len(pend) > 0 {
			for id, monkey := range pend {
				val1, ok := vals[monkey.op1]
				if ok {
					val2, ok := vals[monkey.op2]
					if ok {
						vals[id] = doOperation(val1, val2, monkey.operator)
						delete(pend, id)
						if id == "root" {
							break OUTER
						}
					}
				}
			}
		}
	}

	return vals["root"], -1
}

func doOperation(val1, val2 int, op rune) int {
	if op == SUM {
		return val1 + val2
	} else if op == SUB {
		return val1 - val2
	} else if op == MUL {
		return val1 * val2
	} else if op == DIV {
		return val1 / val2
	} else {
		fmt.Println("Unexpected operator:", op)
		os.Exit(1)
		return -1
	}
}

func readFile(file string) (Values, Pending) {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	vals := make(Values)
	pend := make(Pending)

	for scanner.Scan() {
		line := scanner.Text()
		m := RE_VAL.FindStringSubmatch(line)
		if len(m) > 0 {
			// is a straight value
			val, _ := strconv.Atoi(m[2])
			vals[m[1]] = val
		} else {
			// is an operation
			m = RE_PEND.FindStringSubmatch(line)
			pend[m[1]] = &Monkey{
				op1:      m[2],
				op2:      m[4],
				operator: rune(m[3][0]),
			}
		}
	}

	readFile.Close()
	return vals, pend
}

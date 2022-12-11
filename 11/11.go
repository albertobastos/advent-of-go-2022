package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const ROUNDS_PART1 = 20
const ROUNDS_PART2 = 10000

func reliefPart1(item int) int {
	return item / 3
}
func reliefPart2(item int) int {
	return item
}

type Monkey struct {
	items       []int
	op          func(int) int
	throwTo     func(int) int
	testDivisor int
	inspected   int
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	part1 := doRun(file, ROUNDS_PART1, reliefPart1)
	part2 := doRun(file, ROUNDS_PART2, reliefPart2)
	return part1, part2
}

func doRun(file string, rounds int, relief func(int) int) int {
	monkeys, normalizer := readFile(file)
	for i := 0; i < rounds; i++ {
		doRound(monkeys, relief, normalizer)
	}
	//printState(monkeys)
	return calcMonkeyBusiness(monkeys)
}

func readFile(file string) ([]*Monkey, func(int) int) {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var monkey *Monkey
	list := []*Monkey{}

	for {
		monkey = readMonkey(scanner)
		if monkey == nil {
			break
		}
		list = append(list, monkey)
	}

	normalizedDiv := 1
	for _, m := range list {
		normalizedDiv *= m.testDivisor
	}
	normalizer := func(item int) int {
		return item % normalizedDiv
	}

	readFile.Close()
	return list, normalizer
}

const PFX_ITEMS = len("  Starting items: ")
const PFX_OP = len("  Operation: new = ")
const PFX_TEST = len("  Test: divisible by ")
const PFX_COND_TRUE = len("    If true: throw to monkey ")
const PFX_COND_FALSE = len("    If false: throw to monkey ")

func readMonkey(scanner *bufio.Scanner) *Monkey {
	lines := readMonkeyLines(scanner)
	if lines == nil {
		return nil
	}

	items := []int{}
	for _, itemstr := range strings.Split(lines[1][PFX_ITEMS:], ", ") {
		item, _ := strconv.Atoi(itemstr)
		items = append(items, item)
	}

	operands := strings.Split(lines[2][PFX_OP:], " ")
	op := buildMonkeyOperation(operands)

	throwTo, testDivisor := buildMonkeyThrowTo(lines[3:6])

	return &Monkey{items, op, throwTo, testDivisor, 0}
}

func readMonkeyLines(scanner *bufio.Scanner) []string {
	lines := []string{}
	for i := 0; i < 6 && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}
	scanner.Scan() // skip blank line separator
	if len(lines) < 6 {
		return nil
	} else {
		return lines
	}
}

func buildMonkeyOperation(ops []string) func(int) int {
	if len(ops) != 3 {
		fmt.Println("Unexpected number of operands:", ops)
		os.Exit(1)
	}
	op1_getter := getOperandGetter(ops[0])
	op2_getter := getOperandGetter(ops[2])

	// only sum (+) and multiplication (*) supported
	if ops[1] == "+" {
		return func(old int) int { return op1_getter(old) + op2_getter(old) }
	} else if ops[1] == "*" {
		return func(old int) int { return op1_getter(old) * op2_getter(old) }
	}

	fmt.Println("Unexpected operation type:", ops)
	os.Exit(1)
	return nil
}

func getOperandGetter(opstr string) func(int) int {
	if opstr == "old" {
		return func(x int) int { return x }
	} else {
		op, _ := strconv.Atoi(opstr)
		return func(x int) int { return op }
	}
}

func buildMonkeyThrowTo(lines []string) (func(int) int, int) {
	if len(lines) != 3 {
		fmt.Println("Unexpected number of lines to decide throwing:", lines)
		os.Exit(1)
	}

	testDiv, _ := strconv.Atoi(lines[0][PFX_TEST:])
	monkeyIfTrue, _ := strconv.Atoi(lines[1][PFX_COND_TRUE:])
	monkeyIfFalse, _ := strconv.Atoi(lines[2][PFX_COND_FALSE:])

	return func(new int) int {
		if new%testDiv == 0 {
			return monkeyIfTrue
		} else {
			return monkeyIfFalse
		}
	}, testDiv
}

func doRound(monkeys []*Monkey, relief func(int) int, norm func(int) int) {
	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			item := monkey.op(item)
			item = norm(relief(item))
			monkeyTo := monkeys[monkey.throwTo(item)]
			monkeyTo.items = append(monkeyTo.items, item)
			monkey.inspected++
		}
		monkey.items = []int{}
	}
}

func printState(monkeys []*Monkey) {
	for i, m := range monkeys {
		fmt.Println("Monkey", i, ":", m.inspected, "inspections, items:", m.items)
	}
}

func calcMonkeyBusiness(monkeys []*Monkey) int {
	inspected := []int{}
	for _, m := range monkeys {
		inspected = append(inspected, m.inspected)
	}
	sort.Ints(inspected)
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

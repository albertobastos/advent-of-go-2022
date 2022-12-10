package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DISPLAY_COLS = 40
const DISPLAY_ROWS = 6
const DISPLAY_PIXELS = DISPLAY_COLS * DISPLAY_ROWS

type Machine struct {
	register int
	cycles   int
	display  []rune
}

type Instruction interface {
	execute(*Machine)
	cycles() int
}

type noop struct{}
type addx struct {
	param int
}

func (i noop) execute(m *Machine) {
	// do nothing
}

func (i noop) cycles() int {
	return 1
}

func (i addx) execute(m *Machine) {
	m.register += i.param
}

func (i addx) cycles() int {
	return 2
}

func initMachine() *Machine {
	m := Machine{}
	m.reset()
	return &m
}

func (m *Machine) reset() {
	m.register = 1
	m.cycles = 0
	m.display = []rune{}
	for i := 0; i < DISPLAY_PIXELS; i++ {
		m.display = append(m.display, '.')
	}
}

func (m *Machine) updateDisplay() {
	col := m.cycles % DISPLAY_COLS
	if m.register >= col-1 && m.register <= col+1 {
		m.display[m.cycles] = '#'
	}
}

func (m *Machine) getDisplay() []string {
	d := []string{}
	for row := 0; row < DISPLAY_ROWS; row++ {
		d = append(d, string(m.display[row*DISPLAY_COLS:(row+1)*DISPLAY_COLS]))
	}
	return d
}

func exec(m *Machine, instrs []string) int {
	r := 0
	var op Instruction
	for _, instr := range instrs {
		op = parseInstruction(instr)
		opc := op.cycles()
		for i := 0; i < opc; i++ {
			m.updateDisplay()
			m.cycles++
			if (m.cycles-20)%40 == 0 {
				r += m.cycles * m.register
			}
		}
		op.execute(m)
	}
	return r
}

func parseInstruction(str string) Instruction {
	if str[:4] == "noop" {
		return noop{}
	} else {
		// addx
		param, _ := strconv.Atoi(str[5:])
		return addx{param}
	}
}

func readFile(file string) []string {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	readFile.Close()
	return lines
}

func run(file string) (int, []string) {
	lines := readFile(file)
	machine := initMachine()
	part1 := exec(machine, lines)
	part2 := machine.getDisplay()
	return part1, part2
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:")
	for _, l := range part2 {
		fmt.Println(l)
	}
}

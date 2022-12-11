package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PART1_THRESHOLD = 100000
const TOTAL_DISK = 70000000
const SPACE_REQUIRED = 30000000

type Type string

const CMD_CD = Type("CD")
const CMD_LS = Type("LS")
const DIR = Type("DIR")
const FILE = Type("FILE")

type Line struct {
	t     Type
	extra string // cd, dir or file name
	size  int    // file size
}

type Entry struct {
	name     string
	isDir    bool
	parent   *Entry
	children map[string]*Entry
	size     int
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	root := readFile(file)
	//root.print()
	return doPart1(root), doPart2(root)
}

func readFile(file string) *Entry {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	root := Entry{"/", true, nil, make(map[string]*Entry), -1}
	curr := &root

	for scanner.Scan() {
		l := parseLine(scanner.Text())
		if l.t == CMD_CD && l.extra == ".." {
			curr = curr.parent
		} else if l.t == CMD_CD && l.extra == "/" {
			curr = &root
		} else if l.t == CMD_CD {
			curr = curr.children[l.extra]
		} else if l.t == CMD_LS {
			continue
		} else if l.t == DIR {
			curr.children[l.extra] = l.toDirEntry(curr)
		} else {
			curr.children[l.extra] = l.toFileEntry()
		}
	}

	root.fillSize()

	readFile.Close()
	return &root
}

func parseLine(str string) *Line {
	var l Line
	slen := len(str)
	if slen > 5 && str[0:5] == "$ cd " {
		l = Line{CMD_CD, str[5:], -1}
	} else if str == "$ ls" {
		l = Line{CMD_LS, "", -1}
	} else if slen > 4 && str[0:4] == "dir " {
		l = Line{DIR, str[4:], -1}
	} else {
		spl := strings.Split(str, " ")
		size, _ := strconv.Atoi(spl[0])
		l = Line{FILE, spl[1], size}
	}
	return &l
}

func (l *Line) toDirEntry(parent *Entry) *Entry {
	return &Entry{l.extra, true, parent, make(map[string]*Entry), -1}
}

func (l *Line) toFileEntry() *Entry {
	return &Entry{l.extra, false, nil, nil, l.size}
}

func (e *Entry) fillSize() {
	if !e.isDir {
		return
	}
	sum := 0
	for _, c := range e.children {
		c.fillSize()
		sum += c.size
	}
	e.size = sum
}

func (e *Entry) print() {
	e._print("")
}

func (e *Entry) _print(pfx string) {
	if e.isDir {
		fmt.Println(pfx, "-", e.name, "(dir, size="+fmt.Sprint(e.size)+")")
		npfx := pfx + "  "
		for _, c := range e.children {
			c._print(npfx)
		}
	} else {
		fmt.Println(pfx, "-", e.name, "(file, size="+fmt.Sprint(e.size)+")")
	}
}

func doPart1(e *Entry) int {
	if !e.isDir {
		return 0
	}
	sum := 0
	if e.size < PART1_THRESHOLD {
		sum += e.size
	}
	for _, c := range e.children {
		sum += doPart1(c)
	}
	return sum
}

func doPart2(root *Entry) int {
	free := TOTAL_DISK - root.size
	required := SPACE_REQUIRED - free
	min := root.findSmallestRequired(required)
	if min == nil {
		return -1
	} else {
		return min.size
	}
}

func (e *Entry) findSmallestRequired(r int) *Entry {
	if !e.isDir {
		return nil
	}
	var min *Entry
	if e.size > r {
		min = e
	}
	for _, c := range e.children {
		cmin := c.findSmallestRequired(r)
		if min == nil {
			min = cmin
		} else if cmin != nil && cmin.size < min.size {
			min = cmin
		}
	}
	return min
}

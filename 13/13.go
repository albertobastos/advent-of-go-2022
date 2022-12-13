package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	valueInt  int
	valueList []*Item
}

func (i *Item) isList() bool {
	return i.valueList != nil
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	items := readFile(file)
	l := len(items)
	part1 := 0
	for i := 0; i < l; i = i + 2 {
		if compare(items[i], items[i+1]) <= 0 {
			part1 += (i / 2) + 1
		}
	}

	part2 := 1
	divider1 := &Item{0, []*Item{{0, []*Item{{2, nil}}}}}
	divider2 := &Item{0, []*Item{{0, []*Item{{6, nil}}}}}
	items = append(items, divider1, divider2)
	sort.SliceStable(items, func(i, j int) bool {
		return compare(items[i], items[j]) <= 0
	})
	for i, item := range items {
		if compare(divider1, item) == 0 || compare(divider2, item) == 0 {
			part2 *= (i + 1)
		}
	}
	return part1, part2
}

func readFile(file string) []*Item {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	items := []*Item{}
	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			item, _ := readItem(str)
			items = append(items, item)
		}
	}

	readFile.Close()
	return items
}

func readItem(str string) (*Item, int) {
	idx := 0
	vn := 0
	var vl []*Item = nil

	if str[idx] == '[' {
		vl = []*Item{}
		idx++
		for str[idx] != ']' {
			child, read := readItem(str[idx:])
			vl = append(vl, child)
			idx += read
			if str[idx] == ',' {
				idx++
			}
		}
		idx++
	} else {
		for ; str[idx] >= '0' && str[idx] <= '9'; idx++ {
			vn = (vn * 10) + int(str[idx]) - '0'
		}
	}
	return &Item{vn, vl}, idx
}

func compare(a *Item, b *Item) int {
	if !a.isList() && !b.isList() {
		return a.valueInt - b.valueInt
	} else if a.isList() && b.isList() {
		return compareLists(a.valueList, b.valueList)
	} else {
		if !a.isList() {
			a = &Item{0, []*Item{a}}
		} else if !b.isList() {
			b = &Item{0, []*Item{b}}
		}
		return compare(a, b)
	}
}

func compareLists(a []*Item, b []*Item) int {
	lenb := len(b)
	for i, ca := range a {
		if i >= lenb {
			// right side run out of items, left side is greater
			return 1
		}
		ccmp := compare(ca, b[i])
		if ccmp > 0 {
			return ccmp
		} else if ccmp < 0 {
			return ccmp
		}
	}
	if len(a) == lenb {
		return 0
	} else {
		// left side run out of items, right side is greater
		return -1
	}
}

func (i *Item) toString() string {
	str := ""
	if i.isList() {
		str += "["
		for idx, si := range i.valueList {
			if idx > 0 {
				str += ","
			}
			str += si.toString()
		}
		str += "]"
	} else {
		str += fmt.Sprint(i.valueInt)
	}
	return str
}

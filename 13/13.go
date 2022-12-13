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

type Pair struct {
	first  *Item
	second *Item
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	pairs := readFile(file)

	part1 := 0
	for i, p := range pairs {
		if p.compare() <= 0 {
			part1 += i + 1
		}
	}

	part2 := 1
	divider1 := &Item{0, []*Item{&Item{0, []*Item{&Item{2, nil}}}}}
	divider2 := &Item{0, []*Item{&Item{0, []*Item{&Item{6, nil}}}}}
	allItems := []*Item{divider1, divider2}
	for _, p := range pairs {
		allItems = append(allItems, p.first)
		allItems = append(allItems, p.second)
	}
	sort.SliceStable(allItems, func(i, j int) bool {
		return compare(allItems[i], allItems[j]) <= 0
	})
	for i, item := range allItems {
		if compare(divider1, item) == 0 || compare(divider2, item) == 0 {
			part2 *= (i + 1)
		}
	}
	return part1, part2
}

func readFile(file string) []*Pair {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	pairs := []*Pair{}
	var p *Pair
	for {
		p = readPair(scanner)
		if p == nil {
			break
		} else {
			pairs = append(pairs, p)
		}
	}

	readFile.Close()
	return pairs
}

func readPair(scanner *bufio.Scanner) *Pair {
	if !scanner.Scan() {
		return nil
	}
	first, _ := readItem(scanner.Text())
	scanner.Scan()
	second, _ := readItem(scanner.Text())
	scanner.Scan() // skip blank line separator

	return &Pair{first, second}
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

func (p *Pair) print() {
	fmt.Println(p.first.toString())
	fmt.Println(p.second.toString())
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

func (p *Pair) compare() int {
	return compare(p.first, p.second)
}

func compare(a *Item, b *Item) int {
	if !a.isList() && !b.isList() {
		return compareInts(a.valueInt, b.valueInt)
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

func compareInts(a int, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Comparable interface {
	int8 | int
}

type Tree struct {
	height int8

	mnorth int8
	meast  int8
	msouth int8
	mwest  int8

	score int
}

type Forest [][]*Tree

func initForest(size int) Forest {
	f := make(Forest, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			f[i] = append(f[i], nil)
		}
	}
	return f
}

func (f Forest) initTree(row int, col int, height_str rune) {
	height, _ := strconv.Atoi(string(height_str))
	f[row][col] = &Tree{height: int8(height)}
}

func (f Forest) treeAt(row int, col int) *Tree {
	if row < 0 || col < 0 || row >= len(f) || col >= len(f) {
		return nil
	}
	return f[row][col]
}

func (f Forest) calculateAll() {
	size := len(f)
	for i := 0; i < size; i++ {
		// east -> start at left border and go rightwards
		// south -> start at top border and go downwards
		calculateMaxEast(f, size, i, 0)
		calculateMaxSouth(f, size, 0, i)
	}
	for i := size - 1; i >= 0; i-- {
		// west -> start at right border and go leftwards
		// north -> start at bottom border and go upwards
		calculateMaxWest(f, size, i, size-1)
		calculateMaxNorth(f, size, size-1, i)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			calculateScore(f, size, i, j)
		}
	}
}

func calculateMaxEast(f Forest, size int, x int, y int) int8 {
	if y >= size {
		return -1
	}
	tree := f[x][y]
	tree.meast = calculateMaxEast(f, size, x, y+1)
	return maxint8(tree.height, tree.meast)
}

func calculateMaxWest(f Forest, size int, x int, y int) int8 {
	if y < 0 {
		return -1
	}
	tree := f[x][y]
	tree.mwest = calculateMaxWest(f, size, x, y-1)
	return maxint8(tree.height, tree.mwest)
}

func calculateMaxSouth(f Forest, size int, x int, y int) int8 {
	if x >= size {
		return -1
	}
	tree := f[x][y]
	tree.msouth = calculateMaxSouth(f, size, x+1, y)
	return maxint8(tree.height, tree.msouth)
}

func calculateMaxNorth(f Forest, size int, x int, y int) int8 {
	if x < 0 {
		return -1
	}
	tree := f[x][y]
	tree.mnorth = calculateMaxNorth(f, size, x-1, y)
	return maxint8(tree.height, tree.mnorth)
}

func calculateScore(f Forest, size int, x int, y int) int {
	tree := f[x][y]
	east, south, west, north := 0, 0, 0, 0

	for j := y + 1; j < size; j++ {
		next := f[x][j]
		east++
		if next.height >= tree.height {
			break
		}
	}

	for j := y - 1; j >= 0; j-- {
		next := f[x][j]
		west++
		if next.height >= tree.height {
			break
		}
	}

	for i := x + 1; i < size; i++ {
		next := f[i][y]
		south++
		if next.height >= tree.height {
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		next := f[i][y]
		north++
		if next.height >= tree.height {
			break
		}
	}

	tree.score = east * south * west * north
	return tree.score
}

func maxint8(a int8, b int8) int8 {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxint(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (f Forest) countVisible() int {
	sum := 0
	for _, row := range f {
		for _, tree := range row {
			sum += tree.isVisible()
		}
	}
	return sum
}

func (f Forest) getMaxScore() int {
	max := 0
	for _, row := range f {
		for _, tree := range row {
			max = maxint(max, tree.score)
		}
	}
	return max
}

func (t Tree) isVisible() int {
	if t.height > t.meast || t.height > t.mwest || t.height > t.msouth || t.height > t.mnorth {
		return 1
	} else {
		return 0
	}
}

func (f Forest) print(title string, fn func(t *Tree) int8) {
	fmt.Println(title)
	for _, row := range f {
		for _, col := range row {
			fmt.Printf("%d", fn(col))
		}
		fmt.Printf("\n")
	}
}

func readFile(file string) Forest {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var f Forest
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if f == nil {
			f = initForest(len(line))
		}
		for col, hstr := range line {
			f.initTree(row, col, hstr)
		}
		row++
	}

	readFile.Close()
	return f
}

func run(file string) (int, int) {
	forest := readFile(file)
	forest.calculateAll()
	// forest.print("Heights", func(t *Tree) int8 { return int8(t.height) })
	// forest.print("MaxNorth", func(t *Tree) int8 { return int8(t.mnorth) })
	// forest.print("MaxEast", func(t *Tree) int8 { return int8(t.meast) })
	// forest.print("MaxSouth", func(t *Tree) int8 { return int8(t.vsouth) })
	// forest.print("MaxWest", func(t *Tree) int8 { return int8(t.mwest) })
	// forest.print("Scores", func(t *Tree) int8 { return int8(t.score) })
	part1 := forest.countVisible()
	part2 := forest.getMaxScore()
	return part1, part2
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

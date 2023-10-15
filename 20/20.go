package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

func main() {
	part1, part2 := run("input.txt")
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func run(file string) (int, int) {
	head, zero, size := readFile(file)
	decrypt(head, zero, size)
	return -1, -1
}

func readFile(file string) (*Node, *Node, int) {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var curr *Node = nil
	var zero *Node = nil
	n := 0

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		if curr == nil {
			curr = &Node{val, nil, nil}
			curr.next = curr
			curr.prev = curr
		} else {
			next := &Node{val, curr, curr.next}
			curr.next = next
			curr = next
		}
		if val == 0 {
			zero = curr
		}
		n++
	}

	readFile.Close()
	return curr.next, zero, n
}

func decrypt(head *Node, zero *Node, size int) {
	arr := []*Node{}
	for curr := head; len(arr) < size; curr = curr.next {
		arr = append(arr, curr)
	}

	for _, node := range arr {
		move(node, node.val%size)
	}
}

func move(node *Node, steps int) {
	// TODO
}

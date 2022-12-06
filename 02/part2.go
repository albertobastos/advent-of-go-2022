package main

import (
	"bufio"
	"os"
)

func doPart2(file string) int {
	plays := p2_readFile(file)
	points := 0
	for _, play := range plays {
		points += p2_calculatePoints(play)
	}
	return points
}

func p2_readFile(file string) []Play {
	readFile, _ := os.Open(file)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	plays := []Play{}

	for scanner.Scan() {
		line := scanner.Text()
		plays = append(plays, p2_buildPlay(line))
	}

	readFile.Close()
	return plays
}

func p2_buildPlay(s string) Play {
	return Play{
		opponent: readOpponent(s[0]),
		result:   readResult(s[2]),
	}
}

func readResult(c byte) int {
	if c == 'X' {
		return LOSE
	} else if c == 'Y' {
		return DRAW
	} else if c == 'Z' {
		return WIN
	}
	return INVALID
}

func p2_calculatePoints(p Play) int {
	re := p.result
	op := p.opponent

	if re == DRAW {
		return op + re
	} else if re == LOSE && op == ROCK {
		return SCISSORS + re
	} else if re == LOSE && op == PAPER {
		return ROCK + re
	} else if re == LOSE && op == SCISSORS {
		return PAPER + re
	} else if re == WIN && op == ROCK {
		return PAPER + re
	} else if re == WIN && op == PAPER {
		return SCISSORS + re
	} else if re == WIN && op == SCISSORS {
		return ROCK + re
	}
	return 0
}
